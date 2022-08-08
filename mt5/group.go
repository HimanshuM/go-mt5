package mt5

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/HimanshuM/go-mt5/constants"
)

// GetTotalGroups gets the total groups from MT5
func (m *Client) GetTotalGroups() (int, error) {
	cmd := &Command{
		Command: constants.CmdGroupTotal,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return -1, err
	}
	if !res.Okay() {
		return -1, fmt.Errorf("error getting total groups: %v", res.ReturnValue)
	}
	totalStr, present := res.Parameters[constants.ParamTotal].(string)
	if !present {
		return -1, fmt.Errorf("invalid response for total groups")
	}
	total, err := strconv.Atoi(totalStr)
	if err != nil {
		return -1, fmt.Errorf("unknown number for total groups: %s", totalStr)
	}
	return total, nil
}

// Group struct resembles an MT5 Group object
type Group struct {
	Group                string
	Server               string
	PermissionsFlags     string
	AuthMode             string
	AuthPasswordMin      string
	AuthOTPMode          string
	Company              string
	CompanyPage          string
	CompanyEmail         string
	CompanySupportPage   string
	CompanySupportEmail  string
	CompanyCatalog       string
	Currency             string
	CurrencyDigits       string
	ReportsMode          string
	ReportsFlags         string
	ReportsSMTP          string
	ReportsSMTPLogin     string
	ReportsSMTPPass      string
	NewsMode             string
	NewsCategory         string
	NewsLangs            []string
	MailMode             string
	TradeFlags           string
	TradeTransferMode    string
	TradeInterestrate    string
	TradeVirtualCredit   string
	MarginMode           string
	MarginSOMode         string
	MarginFreeMode       string
	MarginCall           string
	MarginStopOut        string
	MarginFreeProfitMode string
	DemoLeverage         string
	DemoDeposit          string
	LimitHistory         string
	LimitOrders          string
	LimitSymbols         string
	LimitPositions       string
	Commissions          []*Commission
	Symbols              []*GroupSymbol
}

// GetGroupByName gets a group by name
func (m *Client) GetGroupByName(groupName string) (*Group, error) {
	cmd := &Command{
		Command: constants.CmdGroupGet,
		Parameters: map[string]interface{}{
			constants.ParamGroup: groupName,
		},
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error getting group %s: %v", groupName, err)
	}
	var group Group
	if err := json.Unmarshal([]byte(res.Data), &group); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for group get: %v", err)
	}
	return &group, nil
}

// GetGroupByIndex searches for group within MT5 platform
func (m *Client) GetGroupByIndex(index int) (*Group, error) {
	cmd := &Command{
		Command: constants.CmdGroupIndex,
		Parameters: map[string]interface{}{
			constants.ParamIndex: index,
		},
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error getting group by index: %v", res.ReturnValue)
	}
	var group Group
	if err := json.Unmarshal([]byte(res.Data), &group); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for group get: %v", err)
	}
	return &group, nil
}

// GetAllGroups gets all groups
func (m *Client) GetAllGroups() ([]*Group, error) {
	total, err := m.GetTotalGroups()
	if err != nil {
		return nil, err
	}
	groups := make([]*Group, total)
	for i := 0; i < total; i++ {
		group, err := m.GetGroupByIndex(i)
		if err != nil {
			return nil, err
		}
		groups[i] = group
	}
	return groups, nil
}
