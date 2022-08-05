package mt5

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/HimanshuM/go-mt5/constants"
)

type Tick struct {
	Symbol      string
	Digits      string
	Datetime    string
	DatetimeMsc string
	Bid         string
	Ask         string
	Last        string
	Volume      string
	VolumeReal  string
}

type TickStats struct {
	Symbol           string
	Digits           string
	Bid              string
	BidLow           string
	BidHigh          string
	BidDir           string
	Ask              string
	AskLow           string
	AskHigh          string
	AskDir           string
	Last             string
	LastLow          string
	LastHigh         string
	LastDir          string
	Volume           string
	VolumeLow        string
	VolumeHigh       string
	VolumeDir        string
	TradeDeals       string
	TradeVolume      string
	TradeTurnover    string
	TradeInterest    string
	TradeBuyOrders   string
	TradeBuyVolume   string
	TradeSellOrders  string
	TradeSellVolume  string
	Datetime         string
	DatetimeMsc      string
	PriceOpen        string
	PriceClose       string
	PriceAw          string
	PriceChange      string
	PriceVolatility  string
	PriceTheoretical string
	PriceGreeksDelta string
	PriceGreeksTheta string
	PriceGreeksGamma string
	PriceGreeksVega  string
	PriceGreeksRho   string
	PriceGreeksOmega string
	PriceSensitivity string
}

type LastTicks struct {
	TransactionID int
	Ticks         []*Tick
}

type LastTickStats struct {
	TransactionID int
	TickStats     []*TickStats
}

func (m *Client) GetLastTick(id int, symbols ...string) (*LastTicks, error) {
	symbolNames := strings.Join(symbols, ",")
	cmd := &Command{
		Command: constants.CMD_TICK_LAST,
		Parameters: map[string]interface{}{
			constants.PARAM_SYMBOL:        symbolNames,
			constants.PARAM_TICK_TRANS_ID: id,
		},
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error getting last tick: %v", res.ReturnValue)
	}
	ticks := make([]*Tick, 0)
	if err := json.Unmarshal([]byte(res.Data), &ticks); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for last tick: %v", err)
	}
	transID := res.Parameters[constants.PARAM_TICK_TRANS_ID].(string)
	transactionID, err := strconv.Atoi(transID)
	if err != nil {
		return nil, fmt.Errorf("error parsing transaction ID %s of the last tick: %v", transID, err)
	}

	return &LastTicks{
		Ticks:         ticks,
		TransactionID: transactionID,
	}, nil
}

func (m *Client) GetLastTickByGroup(id int, group string, symbols ...string) (*LastTicks, error) {
	symbolNames := strings.Join(symbols, ",")
	cmd := &Command{
		Command: constants.CMD_TICK_LAST_BY_GROUP,
		Parameters: map[string]interface{}{
			constants.PARAM_SYMBOL:        symbolNames,
			constants.PARAM_TICK_TRANS_ID: id,
			constants.PARAM_GROUP:         group,
		},
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error getting last tick: %v", res.ReturnValue)
	}
	ticks := make([]*Tick, 0)
	if err := json.Unmarshal([]byte(res.Data), &ticks); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for last tick: %v", err)
	}
	transID := res.Parameters[constants.PARAM_TICK_TRANS_ID].(string)
	transactionID, err := strconv.Atoi(transID)
	if err != nil {
		return nil, fmt.Errorf("error parsing transaction ID %s of the last tick: %v", transID, err)
	}

	return &LastTicks{
		Ticks:         ticks,
		TransactionID: transactionID,
	}, nil
}

func (m *Client) GetTickStatistics(id int, symbols ...string) (*LastTickStats, error) {
	symbolNames := strings.Join(symbols, ",")
	cmd := &Command{
		Command: constants.CMD_TICK_STATS,
		Parameters: map[string]interface{}{
			constants.PARAM_SYMBOL:        symbolNames,
			constants.PARAM_TICK_TRANS_ID: id,
		},
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error getting last tick: %v", res.ReturnValue)
	}
	tickStats := make([]*TickStats, 0)
	if err := json.Unmarshal([]byte(res.Data), &tickStats); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for last tick: %v", err)
	}
	transID := res.Parameters[constants.PARAM_TICK_TRANS_ID].(string)
	transactionID, err := strconv.Atoi(transID)
	if err != nil {
		return nil, fmt.Errorf("error parsing transaction ID %s of the last tick: %v", transID, err)
	}

	return &LastTickStats{
		TickStats:     tickStats,
		TransactionID: transactionID,
	}, nil
}
