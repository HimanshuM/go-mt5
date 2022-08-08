package mt5

import (
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strconv"

	"github.com/HimanshuM/go-mt5/constants"
	"github.com/sirupsen/logrus"
)

// User struct resembles an MT5 User object
type User struct {
	Login             int     `json:"Login"`
	MainPassword      string  `json:"MainPassword"`
	InvestPassword    string  `json:"InvestPassword"`
	Rights            int     `json:"Rights"`
	Group             string  `json:"Group"`
	Name              string  `json:"Name"`
	Company           string  `json:"Company"`
	Country           string  `json:"Country"`
	Language          int     `json:"Language"`
	City              string  `json:"City"`
	State             string  `json:"State"`
	ZipCode           string  `json:"ZipCode"`
	Address           string  `json:"Address"`
	Phone             string  `json:"Phone"`
	Email             string  `json:"Email"`
	ID                string  `json:"ID"`
	Status            string  `json:"Status"`
	Comment           string  `json:"Comment"`
	Color             int     `json:"Color"`
	PhonePassword     string  `json:"PhonePassword"`
	Leverage          int     `json:"Leverage"`
	Agent             int     `json:"Agent"`
	CertSerialNumber  int     `json:"CertSerialNumber"`
	Registration      int     `json:"Registration"`
	LastAccess        int     `json:"LastAccess"`
	LastIP            string  `json:"LastIP"`
	LastPassChange    int     `json:"LastPassChange"`
	Account           string  `json:"Account"`
	ClientID          int     `json:"ClientID"`
	Balance           float32 `json:"Balance"`
	Credit            float32 `json:"Credit"`
	InterestRate      float32 `json:"InterestRate"`
	CommissionDaily   float32 `json:"CommissionDaily"`
	CommissionMonthly float32 `json:"CommissionMonthly"`
	BalancePrevDay    float32 `json:"BalancePrevDay"`
	BalancePrevMonth  float32 `json:"BalancePrevMonth"`
	EquityPrevDay     float32 `json:"EquityPrevDay"`
	EquityPrevMonth   float32 `json:"EquityPrevMonth"`
	MQID              string  `json:"MQID"`
	TradeAccounts     string  `json:"TradeAccounts"`
	LeadCampaign      string  `json:"LeadCampaign"`
	LeadSource        string  `json:"LeadSource"`
}

// CreateUser creates a user on the MT5 server
func (m *Client) CreateUser(u *User) error {
	body, err := u.toJSON()
	if err != nil {
		return err
	}

	logrus.Infof("body: %s", body)

	cmd := &Command{
		Command:         constants.CmdUserAdd,
		Parameters:      u.constructUserCreateParameters(),
		ResponseHasBody: true,
		Body:            body,
	}
	response, err := m.IssueCommand(cmd)
	if err != nil {
		return err
	}
	if !response.Okay() {
		return fmt.Errorf("error creating user: %v", response.ReturnValue)
	}

	createdUserMap := make(map[string]interface{})
	if err = json.Unmarshal([]byte(response.Data), &createdUserMap); err != nil {
		return fmt.Errorf("error parsing JSON response for created user: %v", err)
	}

	return u.fromJSON(createdUserMap, response.Parameters)
}

// constructUserCreateParameters returns a map created from the MT5User
func (u *User) constructUserCreateParameters() map[string]interface{} {
	return map[string]interface{}{
		constants.ParamUserLogin:        u.Login,
		constants.ParamUserPassMain:     u.MainPassword,
		constants.ParamUserPassInvestor: u.InvestPassword,
		constants.ParamUserRights:       u.Rights,
		constants.ParamUserGroup:        u.Group,
		constants.ParamUserName:         u.Name,
		constants.ParamUserCompany:      u.Company,
		constants.ParamUserLanguage:     u.Country,
		constants.ParamUserCity:         u.City,
		constants.ParamUserState:        u.State,
		constants.ParamUserZipCode:      u.ZipCode,
		constants.ParamUserAddress:      u.Address,
		constants.ParamUserPhone:        u.Phone,
		constants.ParamUserEmail:        u.Email,
		constants.ParamUserID:           u.ID,
		constants.ParamUserStatus:       u.Status,
		constants.ParamUserComment:      u.Comment,
		constants.ParamUserColor:        u.Color,
		constants.ParamUserPassPhone:    u.PhonePassword,
		constants.ParamUserLeverage:     u.Leverage,
		constants.ParamUserAgent:        u.Agent,
	}
}

// toJSON marshalls the MT5User object into JSON
func (u *User) toJSON() (string, error) {
	body, err := json.Marshal(u)
	if err != nil {
		return "", fmt.Errorf("error marshalling user to JSON: %v", err)
	}

	regx := regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)
	return string(regx.ReplaceAllFunc(body, replaceUTF8Marker)), nil
}

// replaceUTF8Marker replaces UTF-8 markers from strings
func replaceUTF8Marker(source []byte) []byte {
	logrus.Infof("source: %s", string(source))
	return []byte("&#" + html.EscapeString(string(source)) + ";")
}

// fromJSON populates Login, Registration, LastAccess and LastPassChange fields from JSON
func (u *User) fromJSON(userMap, parameters map[string]interface{}) error {
	if login, present := parameters[constants.ParamUserLogin]; present {
		if loginInt, err := strconv.Atoi(login.(string)); err != nil {
			return fmt.Errorf("error parsing Login from response parameter: %v", err)
		} else {
			u.Login = loginInt
		}
	} else {
		if login, err := strconv.Atoi(userMap[constants.ParamUserLoginJSON].(string)); err != nil {
			return fmt.Errorf("error parsing Login from response: %v", err)
		} else {
			u.Login = login
		}
	}

	if registration, err := strconv.Atoi(userMap[constants.ParamUserRegitration].(string)); err != nil {
		return fmt.Errorf("error parsing Registration from response: %v", err)
	} else {
		u.Registration = registration
	}
	if lastAccess, err := strconv.Atoi(userMap[constants.ParamUserLastAccess].(string)); err != nil {
		return fmt.Errorf("error parsing Registration from response: %v", err)
	} else {
		u.LastAccess = lastAccess
	}
	if lastPassChange, err := strconv.Atoi(userMap[constants.ParamUserLastPassChange].(string)); err != nil {
		return fmt.Errorf("error parsing Registration from response: %v", err)
	} else {
		u.LastPassChange = lastPassChange
	}
	return nil
}
