package mt5

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

type MT5 struct {
	config       *MT5Config
	conn         *net.TCPConn
	commandCount int
	connected    bool
}

type MT5Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Version  string
	domain   string
}

func (m *MT5) Init(config *MT5Config) {
	m.connected = false
	m.config = config
	m.commandCount = 0
	m.Connect()
	m.Auth()
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

func (m *MT5) IssueCommand(cmd *MT5Command) (*MT5Response, error) {
	m.commandCount++
	if m.commandCount > MAX_COMMANDS {
		m.commandCount = 1
	}
	cmdString, err := ToUTF16LE(cmd.toString())
	logrus.Infof("cmd string (%d): %s", len(cmdString), cmdString)
	if err != nil {
		return nil, err
	}
	format := PACKET_FORMAT
	if cmd.Command == CMD_AUTH_START {
		format = PREFIX_API
	}
	cmdString = fmt.Sprintf(format+"0%s", len(cmdString), m.commandCount, cmdString)
	logrus.Infof("cmd (%d): %s", len(cmdString), cmdString)
	count, err := m.conn.Write([]byte(cmdString))
	if err != nil {
		logrus.Errorf("error writing bytes: %v", err)
		return nil, err
	}
	logrus.Infof("wrote %d bytes", count)
	return m.readResponse(cmd)
}
