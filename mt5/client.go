package mt5

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

// MT5 structure is the base structure for interacting with MT5 server
type MT5 struct {
	config       *MT5Config
	conn         *net.TCPConn
	commandCount int
	connected    bool
	randomCrypt  string
}

// MT5Config structure allows to specify the MT5 server configuration and manager credentials
type MT5Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Version     string
	CryptMethod string
	domain      string
}

// Init initializes the connection with MT5 server and performs auth
func (m *MT5) Init(config *MT5Config) error {
	m.connected = false
	m.config = config
	m.commandCount = 0
	m.Connect()
	if m.config.CryptMethod == "" {
		m.config.CryptMethod = CRYPT_METHOD_DEFAULT
	}
	return m.Auth()
}

func (m *MT5) getDomain() string {
	if m.config.domain != "" {
		return m.config.domain
	}
	m.config.domain = m.config.Host
	if m.config.Port != "" {
		m.config.domain += ":" + m.config.Port
	}
	return m.config.domain
}

// Connect sets up a socket connection with the MT5 server using MT5Config
func (m *MT5) Connect() error {
	remoteAddr, err := net.ResolveTCPAddr("tcp4", m.getDomain())
	if err != nil {
		logrus.Errorf("resolve tcp address error: %v", err)
		return err
	}
	conn, err := net.DialTCP("tcp", nil, remoteAddr)
	if err != nil {
		logrus.Errorf("dial tcp error: %v", err)
		return err
	}
	m.conn = conn
	return nil
}

// IssueCommand sends a command to the MT5 server specified using MT5Command struct
func (m *MT5) IssueCommand(cmd *MT5Command) (*MT5Response, error) {
	logrus.Debugf("executing command: %s", cmd.Command)
	m.commandCount++
	if m.commandCount > MAX_COMMANDS {
		m.commandCount = 1
	}
	cmdString, err := ToUTF16LE(cmd.toString())
	logrus.Debugf("cmd string (%d): %s", len(cmdString), cmdString)
	if err != nil {
		return nil, err
	}
	format := PACKET_FORMAT
	if cmd.Command == CMD_AUTH_START {
		format = PREFIX_API
	}
	cmdString = fmt.Sprintf(format+"0%s", len(cmdString), m.commandCount, cmdString)
	logrus.Debugf("cmd (%d): %s", len(cmdString), cmdString)
	count, err := m.conn.Write([]byte(cmdString))
	if err != nil {
		logrus.Errorf("error writing bytes: %v", err)
		return nil, err
	}
	logrus.Debugf("wrote %d bytes", count)
	return m.readResponse(cmd)
}
