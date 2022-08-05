package mt5

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/HimanshuM/go-mt5/constants"
)

type ServerTimeSetting struct {
	Daylight      bool
	DaylightState bool
	TimeZone      int
	TimeServer    string
	Days          [][]string
}

type ServerTime struct {
	UnixTime  uint
	Timestamp time.Time
}

// GetServerTimeSettings gets the time settings from the MT5 server
func (m *Client) GetServerTimeSettings() (*ServerTimeSetting, error) {
	cmd := &Command{
		Command:         constants.CMD_SERVER_TIME_SETTING,
		ResponseHasBody: true,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if res.IsUnauthorized() {
		return nil, constants.ErrUnauthorized
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error getting server time: %v", res.ReturnValue)
	}
	responseMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(res.Data), &responseMap); err != nil {
		return nil, fmt.Errorf("error parsing JSON response for server time settings: %v", err)
	}
	serverTimeSettings := &ServerTimeSetting{
		TimeServer: responseMap[constants.PARAM_SERVER_TIME_TIMESERVER].(string),
		Days:       responseMap[constants.PARAM_SERVER_TIME_DAYS].([][]string),
	}
	daylight, err := strconv.ParseBool(responseMap[constants.PARAM_SERVER_TIME_DAYLIGHT].(string))
	if err != nil {
		return nil, fmt.Errorf("invalid response %s for Daylight: %v", responseMap[constants.PARAM_SERVER_TIME_DAYLIGHT], err)
	}
	serverTimeSettings.Daylight = daylight

	daylightState, err := strconv.ParseBool(responseMap[constants.PARAM_SERVER_TIME_DAYLIGHTSTATE].(string))
	if err != nil {
		return nil, fmt.Errorf("invalid response %s for DaylightState: %v", responseMap[constants.PARAM_SERVER_TIME_DAYLIGHTSTATE], err)
	}
	serverTimeSettings.DaylightState = daylightState

	tz, err := strconv.Atoi(responseMap[constants.PARAM_SERVER_TIME_TIMEZONE].(string))
	if err != nil {
		return nil, fmt.Errorf("invalid response %s for TimeZone: %v", responseMap[constants.PARAM_SERVER_TIME_TIMEZONE], err)
	}
	serverTimeSettings.TimeZone = tz

	return serverTimeSettings, nil
}

// GetServerTime gets the current time of the MT5 server
func (m *Client) GetServerTime() (*ServerTime, error) {
	timeSettings, err := m.GetServerTimeSettings()
	if err == constants.ErrUnauthorized {
		timeSettings = &ServerTimeSetting{
			TimeZone: 0,
		}
	} else if err != nil {
		return nil, err
	}

	cmd := &Command{
		Command: constants.CMD_SERVER_TIME,
	}
	res, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, err
	}
	if !res.Okay() {
		return nil, fmt.Errorf("error getting server time: %v", res.ReturnValue)
	}
	timeParameter, present := res.Parameters[constants.PARAM_SERVER_TIME]
	if !present {
		return nil, fmt.Errorf("invalid response for server time query")
	}
	components := strings.SplitN(timeParameter.(string), " ", 2)
	if len(components) != 2 {
		return nil, fmt.Errorf("malformed response for server time: %s %+v", timeParameter, len(components))
	}
	unixTime, err := strconv.ParseUint(components[0], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error parsing unix timestamp %s: %v", components[0], err)
	}

	timestampStr := fmt.Sprintf("%s %s", components[1], formatMinutes(timeSettings.TimeZone))
	fmt.Println(timestampStr)
	timestamp, err := time.Parse("2006.01.02 15:04:05 -07:00", timestampStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing timestamp %s: %v", components[1], err)
	}
	return &ServerTime{
		UnixTime:  uint(unixTime),
		Timestamp: timestamp,
	}, nil
}

func formatMinutes(minutes int) string {
	hours := int(minutes / 60)
	if minutes < 0 {
		minutes = -minutes
	}
	minutes = int(minutes % 60)
	return fmt.Sprintf("%+03d:%02d", hours, minutes)
}
