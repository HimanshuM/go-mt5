package mt5

import (
	"fmt"

	"github.com/HimanshuM/go-mt5/constants"
)

func (m *Client) SendMarketBuyOrderRequest(trade *Trade) error {
	trade.Type = constants.ParamTradeBuyOrder
	return m.SendMarketOrderRequest(trade)
}

func (m *Client) SendMarketSellOrderRequest(trade *Trade) error {
	trade.Type = constants.ParamTradeSellOrder
	return m.SendMarketOrderRequest(trade)
}

func (m *Client) SendMarketOrderRequest(trade *Trade) error {
	trade.Action = constants.ParamTradeMarketOrder
	return m.SendOrderRequest(trade)
}

func (m *Client) SendPositionModifyRequest(trade *Trade) error {
	trade.Action = constants.ParamTradeModifyPosition
	return m.SendOrderRequest(trade)
}

func (m *Client) SendOrderModifyRequest(trade *Trade) error {
	trade.Action = constants.ParamTradeModifyOrder
	return m.SendOrderRequest(trade)
}

func (m *Client) SendOrderCancelRequest(trade *Trade) error {
	trade.Action = constants.ParamTradeCancelOrder
	return m.SendOrderRequest(trade)
}

func (m *Client) SendOrderRequest(trade *Trade) error {
	body, err := trade.toJSON()
	if err != nil {
		return err
	}
	cmd := &Command{
		Command:         constants.CmdTradeRequest,
		Body:            body,
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return err
	}
	if !res.Okay() {
		return fmt.Errorf("error sending trade request: %v", res.ReturnValue)
	}
	return trade.consumeResponse(res)
}

func (m *Client) GetTradeResult(trade *Trade) (*TradeResult, error) {
	cmd := &Command{
		Command: constants.CmdTradeResult,
		Parameters: map[string]interface{}{
			constants.ParamID: trade.ID,
		},
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error sending trade request: %v", res.ReturnValue)
	}
	return consumeTradeResponse(res, trade.ID)
}
