package mt5

import (
	"fmt"

	"github.com/HimanshuM/go-mt5/constants"
)

// Balance structure for all kinds of balance trades
type Balance struct {
	Login       string
	Amount      int64
	Comment     string
	Ticket      string
	CheckMargin bool
}

// SetBalance performs deposit/withdraw actions
func (m *Client) SetBalance(t *Balance) error {
	checkMargin := 0
	if t.CheckMargin {
		checkMargin = 1
	}
	cmd := &Command{
		Command: constants.CmdTradeBalance,
		Parameters: map[string]interface{}{
			constants.ParamTradeLogin:       t.Login,
			constants.ParamTradeType:        constants.ConstTradeBalance,
			constants.ParamTradeBalance:     t.Amount,
			constants.ParamTradeComment:     t.Comment,
			constants.ParamTradeCheckMargin: checkMargin,
		},
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return err
	}
	if !res.Okay() {
		return fmt.Errorf("error setting balance: %v", res.ReturnValue)
	}
	if ticket, present := res.Parameters[constants.ParamTradeTicket]; present {
		t.Ticket = ticket.(string)
	}
	return nil
}
