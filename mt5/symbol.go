package mt5

import (
	"encoding/json"
	"fmt"
)

type Session struct {
	Open  string
	Close string
}

type Symbol struct {
	Symbol                         string
	Path                           string
	ISIN                           string
	CFI                            string
	Category                       string
	Exchange                       string
	Description                    string
	International                  string
	Sector                         string
	Industry                       string
	Country                        string
	Basis                          string
	Source                         string
	Page                           string
	CurrencyBase                   string
	CurrencyBaseDigits             string
	CurrencyProfit                 string
	CurrencyProfitDigits           string
	CurrencyMargin                 string
	CurrencyMarginDigits           string
	Color                          string
	ColorBackground                string
	Digits                         string
	Point                          string
	Multiply                       string
	TickFlags                      string
	TickBookDepth                  string
	FilterSoft                     string
	FilterSoftTicks                string
	FilterHard                     string
	FilterHardTicks                string
	FilterDiscard                  string
	FilterSpreadMax                string
	FilterSpreadMin                string
	FilterGapTicks                 string
	SubscriptionsDelay             string
	TickChartMode                  string
	TradeMode                      string
	CalcMode                       string
	ExecMode                       string
	GTCMode                        string
	FillFlags                      string
	ExpirFlags                     string
	Spread                         string
	SpreadBalance                  string
	SpreadDiff                     string
	SpreadDiffBalance              string
	TickValue                      string
	TickSize                       string
	ContractSize                   string
	StopsLevel                     string
	FreezeLevel                    string
	QuotesTimeout                  string
	VolumeMin                      string
	VolumeMax                      string
	VolumeStep                     string
	VolumeLimit                    string
	MarginFlags                    string
	MarginInitial                  string
	MarginMaintenance              string
	MarginInitialBuy               string
	MarginInitialSell              string
	MarginInitialBuyLimit          string
	MarginInitialSellLimit         string
	MarginInitialBuyStop           string
	MarginInitialSellStop          string
	MarginInitialBuyStopLimit      string
	MarginInitialSellStopLimit     string
	MarginMaintenanceBuy           string
	MarginMaintenanceSell          string
	MarginMaintenanceBuyLimit      string
	MarginMaintenanceSellLimit     string
	MarginMaintenanceBuyStop       string
	MarginMaintenanceSellStop      string
	MarginMaintenanceBuyStopLimit  string
	MarginMaintenanceSellStopLimit string
	MarginHedged                   string
	MarginRateCurrency             string
	MarginRateLiquidity            string
	SwapMode                       string
	SwapLong                       string
	SwapShort                      string
	Swap3Day                       string
	SwapYearDays                   string
	SwapFlags                      string
	SwapRateSunday                 string
	SwapRateMonday                 string
	SwapRateTuesday                string
	SwapRateWednesday              string
	SwapRateThursday               string
	SwapRateFriday                 string
	SwapRateSaturday               string
	TimeStart                      string
	TimeExpiration                 string
	SessionsQuotes                 [][]Session
	SessionsTrades                 [][]Session
	REFlags                        string
	RETimeout                      string
	IECheckMode                    string
	IETimeout                      string
	IESlipProfit                   string
	IESlipLosing                   string
	IEVolumeMax                    string
	PriceSettle                    string
	PriceLimitMax                  string
	PriceLimitMin                  string
	TradeFlags                     string
	OrderFlags                     string
	FaceValue                      string
	AccruedInterest                string
	SpliceType                     string
	SpliceTimeType                 string
	SpliceTimeDays                 string
	ChartMode                      string
	OptionMode                     string
	PriceStrike                    string
}

// GetAllSymbols returns all available symbols
func (m *MT5) GetAllSymbols() ([]string, error) {
	cmd := &MT5Command{
		Command:         CMD_SYMBOL_LIST,
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error setting balance: %v", res.ReturnValue)
	}
	symbols := make([]string, 0)
	if err := json.Unmarshal([]byte(res.Data), &symbols); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for symbol get: %v", err)
	}
	return symbols, nil
}

// SearchSymbols searches for symbols with MT5 platform
func (m *MT5) GetSymbol(symbolStr string) (*Symbol, error) {
	cmd := &MT5Command{
		Command: CMD_SYMBOL_GET,
		Parameters: map[string]interface{}{
			PARAM_SYMBOL: symbolStr,
		},
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error setting balance: %v", res.ReturnValue)
	}
	var symbol Symbol
	if err := json.Unmarshal([]byte(res.Data), &symbol); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for symbol get: %v", err)
	}
	return &symbol, nil
}

// GetSymbolByIndex searches for symbols with MT5 platform
func (m *MT5) GetSymbolByIndex(index uint) (*Symbol, error) {
	cmd := &MT5Command{
		Command: CMD_SYMBOL_NEXT,
		Parameters: map[string]interface{}{
			PARAM_SYMBOL_INDEX: index,
		},
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error setting balance: %v", res.ReturnValue)
	}
	var symbol Symbol
	if err := json.Unmarshal([]byte(res.Data), &symbol); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for symbol get: %v", err)
	}
	return &symbol, nil
}
