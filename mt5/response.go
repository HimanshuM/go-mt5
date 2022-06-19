package mt5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type MT5Response struct {
	BodySize     int
	CommandCount int
	Flag         int
	CommandName  string
	ReturnCode   int
	ReturnValue  string
	Parameters   map[string]interface{}
	Data         string
}

func (c *MT5Command) parseResponse(response string, hasBody bool) (*MT5Response, error) {
	meta := response[0:META_SIZE]
	bodySize, err := strconv.ParseInt(meta[0:4], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("error decoding body size from response: %v", err)
	}
	logrus.Infof("body size: %d", int(bodySize))
	cmdCount, err := strconv.ParseInt(meta[4:8], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("error decoding command count from response: %v", err)
	}
	flag, err := strconv.ParseInt(meta[8:9], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error decoding flag from response: %v", err)
	}
	headerLength := bodySize + 1
	if hasBody {
		headerLength := strings.Index(response[META_SIZE:], "\n")
		if headerLength < 0 {
			return nil, fmt.Errorf("malformed response header: %v", response[META_SIZE:])
		}
	}
	header := response[META_SIZE:headerLength]
	logrus.Infof("header (%d): %s", len(header), header)
	if header, err = ToUTF8(header); err != nil {
		return nil, fmt.Errorf("error converting header to UTF-8: %v", err)
	}
	logrus.Infof("header (utf-8) (%d): %s", len(header), header)
	headerComponents := strings.Split(header, "|")
	logrus.Infof("components: %v", headerComponents)

	mt5Response := &MT5Response{
		BodySize:     int(bodySize),
		CommandCount: int(cmdCount),
		Flag:         int(flag),
		CommandName:  headerComponents[0],
		Data:         response[headerLength:],
	}
	mt5Response.parseParameters(headerComponents[1:])
	return mt5Response, nil
}

func (mt5Response *MT5Response) parseParameters(components []string) error {
	mt5Response.Parameters = make(map[string]interface{})
	for _, parameter := range components {
		if len(parameter) == 0 {
			continue
		}
		components := strings.SplitN(parameter, "=", 2)
		if components[0] == PARAM_RETURN_CODE {
			err := mt5Response.parseReturnString(components[1])
			if err != nil {
				return err
			}
		} else {
			mt5Response.Parameters[components[0]] = components[1]
		}
	}
	return nil
}

func (mt5Response *MT5Response) parseReturnString(returnStr string) error {
	logrus.Infof("return str (%d): %s", len(returnStr), returnStr)
	components := strings.SplitN(returnStr, " ", 2)
	retCode, err := strconv.ParseInt(components[0], 10, 32)
	if err != nil {
		return fmt.Errorf("error parsing return parameter: %v", err)
	}
	mt5Response.ReturnCode = int(retCode)
	mt5Response.ReturnValue = components[1]
	return nil
}
