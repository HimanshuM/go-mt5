package mt5

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/HimanshuM/go-mt5/constants"
	"github.com/HimanshuM/go-mt5/encoding"
	"github.com/sirupsen/logrus"
)

// Response struct holds the response received from the MT5 server against a command
type Response struct {
	BodySize     int
	CommandCount int
	Flag         int
	CommandName  string
	ReturnCode   int
	ReturnValue  string
	Parameters   map[string]interface{}
	Data         string
}

// Okay checks if the response was a success (RETCODE = 0 or 1)
func (res *Response) Okay() bool {
	return res.ReturnCode == 0 || res.ReturnCode == 1
}

func (res *Response) IsUnauthorized() bool {
	return res.ReturnCode == 8
}

// readResponse reads the response from the socket connection and builds the MT5Response object
func (m *Client) readResponse(cmd *Command) (*Response, error) {
	if cmd.Command == constants.CmdClose {
		return nil, nil
	}
	bufferMeta := new(bytes.Buffer)
	bytesRead, err := io.CopyN(bufferMeta, m.conn, constants.MetaSize)
	if err != nil || bytesRead != constants.MetaSize {
		return nil, fmt.Errorf("invalid response received: %s", bufferMeta.String())
	}
	response, err := parseMeta(bufferMeta.String())
	if err != nil {
		return nil, err
	}
	if response.BodySize == 0 && response.CommandCount > constants.MaxCommands {
		return nil, constants.ErrDisconnected
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
	logrus.Debugf("response: %s", responseStr)
	headerLength := response.BodySize
	if cmd.ResponseHasBody {
		headerLength = strings.Index(responseStr, "\n")
	}
	header := responseStr[:headerLength]
	if header, err = encoding.ToUTF8(header); err != nil {
		return nil, fmt.Errorf("error converting header to UTF-8: %v", err)
	}
	headerComponents := strings.Split(header, "|")
	response.CommandName = headerComponents[0]
	if err := response.parseParameters(headerComponents[1:]); err != nil {
		return nil, err
	}

	body, err := encoding.ToUTF8(responseStr[headerLength:])
	if err != nil {
		return nil, fmt.Errorf("error converting body to UTF-8: %v", err)
	}
	response.Data = body
	return response, nil
}

// parseMeta parses the initial 9 bytes of the response that help parse the response
func parseMeta(response string) (*Response, error) {
	meta := response[0:constants.MetaSize]
	bodySize, err := strconv.ParseInt(meta[0:4], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("error decoding body size from response: %v", err)
	}
	cmdCount, err := strconv.ParseInt(meta[4:8], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("error decoding command count from response: %v", err)
	}
	flag, err := strconv.Atoi(meta[8:9])
	if err != nil {
		return nil, fmt.Errorf("error decoding flag from response: %v", err)
	}

	return &Response{
		BodySize:     int(bodySize),
		CommandCount: int(cmdCount),
		Flag:         flag,
	}, nil
}

// parseParameters constructs a map from the response parameters
func (res *Response) parseParameters(components []string) error {
	res.Parameters = make(map[string]interface{})
	for _, parameter := range components {
		parameter = strings.Trim(parameter, "\r\n ")
		if parameter == "" {
			continue
		}
		components := strings.SplitN(parameter, "=", 2)
		if len(components) <= 1 {
			continue
		}
		if components[0] == constants.ParamReturnCode {
			err := res.parseReturnString(components[1])
			if err != nil {
				return err
			}
		} else {
			res.Parameters[components[0]] = components[1]
		}
	}
	return nil
}

// parseRetunString gets the return code and description from the RETCODE parameter
func (res *Response) parseReturnString(returnStr string) error {
	components := strings.SplitN(returnStr, " ", 2)
	retCode, err := strconv.ParseInt(components[0], 10, 32)
	if err != nil {
		return fmt.Errorf("error parsing return parameter: %v", err)
	}
	res.ReturnCode = int(retCode)
	res.ReturnValue = components[1]
	return nil
}
