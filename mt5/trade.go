package mt5

import (
	"fmt"

	"github.com/HimanshuM/go-mt5/constants"
)

// Trade structure for all kinds of trades
type Trade struct {
	Login       string
	Amount      int64
	Comment     string
	Ticket      string
	CheckMargin bool
}

// SetBalance performs deposit/withdraw actions
func (m *Client) SetBalance(t *Trade) error {
	checkMargin := 0
	if t.CheckMargin {
		checkMargin = 1
	}
	cmd := &Command{
		Command: constants.CMD_TRADE_BALANCE,
		Parameters: map[string]interface{}{
			constants.PARAM_TRADE_LOGIN:        t.Login,
			constants.PARAM_TRADE_TYPE:         constants.CONST_TRADE_BALANCE,
			constants.PARAM_TRADE_BALANCE:      t.Amount,
			constants.PARAM_TRADE_COMMENT:      t.Comment,
			constants.PARAM_TRADE_CHECK_MARGIN: checkMargin,
		},
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return err
	}
	if !res.Okay() {
		return fmt.Errorf("error setting balance: %v", res.ReturnValue)
	}
	if ticket, present := res.Parameters[constants.PARAM_TRADE_TICKET]; present {
		t.Ticket = ticket.(string)
	}
	return nil
}
