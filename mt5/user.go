package mt5

import (
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

// MT5User struct resembles an MT5 MT5User object
type MT5User struct {
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

func (m *MT5) CreateUser(u *MT5User) error {
	body, err := u.toJSON()
	if err != nil {
		return err
	}

	logrus.Infof("body: %s", string(body))

	cmd := &MT5Command{
		Command:         CMD_USER_ADD,
		Parameters:      u.constructUserCreateParameters(),
		ResponseHasBody: true,
		Body:            string(body),
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

func (u *MT5User) constructUserCreateParameters() map[string]interface{} {
	return map[string]interface{}{
		PARAM_USER_LOGIN:         u.Login,
		PARAM_USER_PASS_MAIN:     u.MainPassword,
		PARAM_USER_PASS_INVESTOR: u.InvestPassword,
		PARAM_USER_RIGHTS:        u.Rights,
		PARAM_USER_GROUP:         u.Group,
		PARAM_USER_NAME:          u.Name,
		PARAM_USER_COMPANY:       u.Company,
		PARAM_USER_LANGUAGE:      u.Country,
		PARAM_USER_CITY:          u.City,
		PARAM_USER_STATE:         u.State,
		PARAM_USER_ZIPCODE:       u.ZipCode,
		PARAM_USER_ADDRESS:       u.Address,
		PARAM_USER_PHONE:         u.Phone,
		PARAM_USER_EMAIL:         u.Email,
		PARAM_USER_ID:            u.ID,
		PARAM_USER_STATUS:        u.Status,
		PARAM_USER_COMMENT:       u.Comment,
		PARAM_USER_COLOR:         u.Color,
		PARAM_USER_PASS_PHONE:    u.PhonePassword,
		PARAM_USER_LEVERAGE:      u.Leverage,
		PARAM_USER_AGENT:         u.Agent,
	}
}

func (u *MT5User) toJSON() (string, error) {
	body, err := json.Marshal(u)
	if err != nil {
		return "", fmt.Errorf("error marshalling user to JSON: %v", err)
	}

	regx, err := regexp.Compile(`\\u([0-9a-fA-F]{4})`)
	if err != nil {
		return "", fmt.Errorf("error compiling regular expression for user creation: %v", err)
	}
	return string(regx.ReplaceAllFunc(body, replaceUTF8Marker)), nil
}

func replaceUTF8Marker(source []byte) []byte {
	logrus.Infof("source: %s", string(source))
	return []byte("&#" + html.EscapeString(string(source)) + ";")
}

func (u *MT5User) fromJSON(userMap map[string]interface{}, parameters map[string]interface{}) error {
	if login, present := parameters[PARAM_USER_LOGIN]; present {
		if loginInt, err := strconv.Atoi(login.(string)); err != nil {
			return fmt.Errorf("error parsing Login from response parameter: %v", err)
		} else {
			u.Login = loginInt
		}
	} else {
		if login, err := strconv.Atoi(userMap[PARAM_USER_LOGIN_JSON].(string)); err != nil {
			return fmt.Errorf("error parsing Login from response: %v", err)
		} else {
			u.Login = int(login)
		}
	}

	if registration, err := strconv.Atoi(userMap[PARAM_USER_REGISTRATION].(string)); err != nil {
		return fmt.Errorf("error parsing Registration from response: %v", err)
	} else {
		u.Registration = int(registration)
	}
	if lastAccess, err := strconv.Atoi(userMap[PARAM_USER_LAST_ACCESS].(string)); err != nil {
		return fmt.Errorf("error parsing Registration from response: %v", err)
	} else {
		u.LastAccess = int(lastAccess)
	}
	if lastPassChange, err := strconv.Atoi(userMap[PARAM_USER_LAST_PASS_CHANGE].(string)); err != nil {
		return fmt.Errorf("error parsing Registration from response: %v", err)
	} else {
		u.LastPassChange = int(lastPassChange)
	}
	return nil
}
