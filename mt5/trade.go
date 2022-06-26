package mt5

import (
	"fmt"
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
func (m *MT5) SetBalance(t *Trade) error {
	checkMargin := 0
	if t.CheckMargin {
		checkMargin = 1
	}
	cmd := &MT5Command{
		Command: CMD_TRADE_BALANCE,
		Parameters: map[string]interface{}{
			PARAM_TRADE_LOGIN:        t.Login,
			PARAM_TRADE_TYPE:         CONST_TRADE_BALANCE,
			PARAM_TRADE_BALANCE:      t.Amount,
			PARAM_TRADE_COMMENT:      t.Comment,
			PARAM_TRADE_CHECK_MARGIN: checkMargin,
		},
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return err
	}
	if !res.Okay() {
		return fmt.Errorf("error setting balance: %v", res.ReturnValue)
	}
	if ticket, present := res.Parameters[PARAM_TRADE_TICKET]; present {
		t.Ticket = ticket.(string)
	}
	return nil
}
