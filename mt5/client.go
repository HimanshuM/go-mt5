package mt5

import (
	"fmt"
	"net"

	"github.com/HimanshuM/go-mt5/constants"
	"github.com/HimanshuM/go-mt5/encoding"
	"github.com/sirupsen/logrus"
)

// Client structure is the base structure for interacting with Client server
type Client struct {
	config       *Config
	conn         *net.TCPConn
	commandCount int
	connected    bool
	randomCrypt  string
}

// Config structure allows to specify the MT5 server configuration and manager credentials
type Config struct {
	Host          string
	Port          string
	Username      string
	Password      string
	Version       string
	CryptMethod   string
	RetryAttempts int
	domain        string
}

// Init initializes the connection with MT5 server and performs auth
func (m *Client) Init(config *Config) error {
	m.connected = false
	m.config = config
	m.commandCount = 0
	if m.config.CryptMethod == "" {
		m.config.CryptMethod = constants.CRYPT_METHOD_DEFAULT
	}
	if config.RetryAttempts <= 0 {
		config.RetryAttempts = 3
	} else if config.RetryAttempts > 5 {
		logrus.Warnf("Too many retries defined (%d) for connection errors, 3 retries recommended", config.RetryAttempts)
	}
	return m.Reconnect()
}

func (m *Client) getDomain() string {
	if m.config.domain != "" {
		return m.config.domain
	}
	m.config.domain = m.config.Host
	if m.config.Port != "" {
		m.config.domain += ":" + m.config.Port
	}
	return m.config.domain
}

// Reconnect sets up a fresh connection to MT5 server and initiates Auth procedure
func (m *Client) Reconnect() error {
	if err := m.Connect(); err != nil {
		return err
	}
	return m.Auth()
}

// Connect sets up a socket connection with the MT5 server using MT5Config
func (m *Client) Connect() error {
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

// Close closes connection to MT5 server
func (m *Client) Close() error {
	cmd := &Command{
		Command: constants.CMD_CLOSE,
	}
	if _, err := m.retryCommand(cmd); err != nil {
		logrus.Infof("Quit command returned error: %v", err)
	}
	return m.conn.Close()
}

// IssueCommand sends a command to the MT5 server specified using MT5Command struct
func (m *Client) IssueCommand(cmd *Command) (*Response, error) {
	for retries := 1; retries <= m.config.RetryAttempts; retries++ {
		response, err := m.retryCommand(cmd)
		if err == constants.ErrDisconnected {
			logrus.Warnf("Connection lost, reconnecting and retrying. Attempt: %d", retries)
			m.connected = false
			if err := m.Reconnect(); err != nil {
				return nil, constants.ErrDisconnected
			}
		} else {
			return response, err
		}
	}
	logrus.Errorf("Failed to reconnected %d times, stopping retries.", m.config.RetryAttempts)
	return nil, constants.ErrDisconnected
}

func (m *Client) retryCommand(cmd *Command) (*Response, error) {
	logrus.Debugf("executing command: %s", cmd.Command)
	m.commandCount++
	if m.commandCount > constants.MAX_COMMANDS {
		m.commandCount = 1
	}
	cmdString, err := encoding.ToUTF16LE(cmd.toString())
	logrus.Debugf("cmd string (%d): %s", len(cmdString), cmdString)
	if err != nil {
		return nil, err
	}
	format := constants.PACKET_FORMAT
	if cmd.Command == constants.CMD_AUTH_START {
		format = constants.PREFIX_API
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
