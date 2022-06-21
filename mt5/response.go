package mt5

import (
	"bytes"
	"fmt"
	"io"
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

func (res *MT5Response) Okay() bool {
	return res.ReturnCode == 0 || res.ReturnCode == 1
}

func (m *MT5) readResponse(cmd *MT5Command) (*MT5Response, error) {
	bufferMeta := new(bytes.Buffer)
	bytesRead, err := io.CopyN(bufferMeta, m.conn, META_SIZE)
	if err != nil || bytesRead != META_SIZE {
		return nil, fmt.Errorf("invalid response received: %s", bufferMeta.String())
	}
	response, err := parseMeta(bufferMeta.String())
	if err != nil {
		return nil, err
	}
	responseStr := ""
	for readBytes := 0; readBytes < response.BodySize; {
		bufferResponse := new(strings.Builder)
		bytesRead, err := io.CopyN(bufferResponse, m.conn, int64(response.BodySize)-int64(readBytes))
		if err != nil {
			return nil, fmt.Errorf("error reading response from socket: %v", err)
		}
		responseStr += bufferResponse.String()
		readBytes += int(bytesRead)
	}
	logrus.Infof("response: %s", responseStr)
	headerLength := response.BodySize
	if cmd.ResponseHasBody {
		headerLength = strings.Index(responseStr, "\n")
	}
	header := responseStr[:headerLength]
	if header, err = ToUTF8(header); err != nil {
		return nil, fmt.Errorf("error converting header to UTF-8: %v", err)
	}
	headerComponents := strings.Split(header, "|")
	response.CommandName = headerComponents[0]
	response.parseParameters(headerComponents[1:])
	return response, nil
}

func parseMeta(response string) (*MT5Response, error) {
	meta := response[0:META_SIZE]
	bodySize, err := strconv.ParseInt(meta[0:4], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("error decoding body size from response: %v", err)
	}
	cmdCount, err := strconv.ParseInt(meta[4:8], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("error decoding command count from response: %v", err)
	}
	flag, err := strconv.ParseInt(meta[8:9], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error decoding flag from response: %v", err)
	}

	return &MT5Response{
		BodySize:     int(bodySize),
		CommandCount: int(cmdCount),
		Flag:         int(flag),
	}, nil
}

func (mt5Response *MT5Response) parseParameters(components []string) error {
	mt5Response.Parameters = make(map[string]interface{})
	for _, parameter := range components {
		parameter = strings.Trim(parameter, "\r\n ")
		if len(parameter) == 0 {
			continue
		}
		components := strings.SplitN(parameter, "=", 2)
		if len(components) <= 1 {
			continue
		}
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
	components := strings.SplitN(returnStr, " ", 2)
	retCode, err := strconv.ParseInt(components[0], 10, 32)
	if err != nil {
		return fmt.Errorf("error parsing return parameter: %v", err)
	}
	mt5Response.ReturnCode = int(retCode)
	mt5Response.ReturnValue = components[1]
	return nil
}
