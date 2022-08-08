package mt5

import (
	"github.com/HimanshuM/go-mt5/constants"
)

// Ping pings the MT5 server to check connection status.
// It is recommended to send this command every 20sec for connections that stay idle
func (m *Client) Ping() error {
	cmd := &Command{
		Command: constants.CmdPing,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return err
	}
	if !res.Okay() {
		return constants.ErrDisconnected
	}
	return nil
}
