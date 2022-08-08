package mt5

import (
	"crypto/md5"
	"fmt"

	"github.com/HimanshuM/go-mt5/constants"
	"github.com/HimanshuM/go-mt5/encoding"
)

// Auth performs authorization with the MT5 server
func (m *Client) Auth() error {
	if m.connected {
		return nil
	}

	m.commandCount = 0

	if !m.config.isEncryptMethodKnown() {
		return fmt.Errorf("unknown encryption method: %s", m.config.CryptMethod)
	}
	resAuthStart, err := m.sendAuthStart()
	if err != nil {
		return err
	}
	srvRand, found := resAuthStart.Parameters[constants.ParamAuthSrvRand]
	if !found {
		return fmt.Errorf("response param %s not found in response", constants.ParamAuthSrvRand)
	}
	passwordHash, err := m.getAuthHash(srvRand.(string))
	if err != nil {
		return err
	}
	randomHex := getRandomHex(16)
	resAuthAnswer, err := m.sendAuthAnswer(passwordHash, randomHex)
	if err != nil {
		return err
	}
	validResponse, err := m.validateAuthAnswer(resAuthAnswer, randomHex)
	if err != nil {
		return err
	}
	if validResponse {
		randomCrypt, present := resAuthAnswer.Parameters[constants.ParamAuthCryptRand]
		if !present {
			return fmt.Errorf("auth answer response does not contain %s", constants.ParamAuthCryptRand)
		}
		m.randomCrypt = randomCrypt.(string)
	}
	m.connected = validResponse
	return nil
}

// isEncryptMethodKnown checks if the encryption method is implemented
func (c *Config) isEncryptMethodKnown() bool {
	return c.CryptMethod == constants.CryptMethodDefault || c.CryptMethod == constants.CryptMethodNone
}

// sendAuthStart sends the AUTH_START command to the MT5 server
func (m *Client) sendAuthStart() (*Response, error) {
	cmd := &Command{
		Command: constants.CmdAuthStart,
		Parameters: map[string]interface{}{
			"VERSION":      constants.APIVersion,
			"AGENT":        constants.WordAPI,
			"LOGIN":        m.config.Username,
			"TYPE":         constants.WordManager,
			"CRYPT_METHOD": m.config.CryptMethod,
		},
	}
	response, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, fmt.Errorf("Auth failed at %s: %v", constants.CmdAuthStart, err)
	}
	if response.CommandName != constants.CmdAuthStart {
		return nil, fmt.Errorf("response of %s (%d) is invalid: %s (%d)", constants.CmdAuthStart, len(constants.CmdAuthStart), response.CommandName, len(response.CommandName))
	}
	if response.ReturnCode != 0 {
		return nil, fmt.Errorf("authorization failed: %v", response.ReturnValue)
	}
	return response, nil
}

// sendAuthAnswer sends AUTH_ANSWER command to the MT5 server
func (m *Client) sendAuthAnswer(passwordHash, randomHex string) (*Response, error) {
	cmd := &Command{
		Command: constants.CmdAuthAnswer,
		Parameters: map[string]interface{}{
			constants.ParamAuthSrvRandAnswer: passwordHash,
			constants.ParamAuthCLIRand:       randomHex,
		},
	}
	response, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, fmt.Errorf("auth failed at %s: %v", constants.CmdAuthAnswer, err)
	}
	if response.CommandName != constants.CmdAuthAnswer {
		return nil, fmt.Errorf("response of %s (%d) is invalid: %s (%d)", constants.CmdAuthAnswer, len(constants.CmdAuthAnswer), response.CommandName, len(response.CommandName))
	}
	if response.ReturnCode != 0 {
		return nil, fmt.Errorf("authorization failed: %v", response.ReturnValue)
	}
	return response, nil
}

// getAuthHash returns an MD5 hash of the password with a given hex string
func (m *Client) getAuthHash(hexString string) (string, error) {
	utf16LEPassword, err := encoding.ToUTF16LE(m.config.Password)
	if err != nil {
		return "", err
	}
	passwordHash := md5.Sum([]byte(utf16LEPassword))
	saltedPassword := string(passwordHash[:]) + constants.WordAPI
	saltedPasswordHash := md5.Sum([]byte(saltedPassword))
	parsedHexString, err := parseHexString(hexString)
	if err != nil {
		return "", err
	}
	finalString := string(saltedPasswordHash[:]) + parsedHexString
	finalHash := md5.Sum([]byte(finalString))
	finalHashHex := ""
	for _, each := range finalHash {
		finalHashHex += fmt.Sprintf("%02x", each)
	}
	return finalHashHex, nil
}

// validateAuthAnswer validates the CLI_RAND_ANSWER against the password hash using CLI_RAND
func (m *Client) validateAuthAnswer(resAuthAnswer *Response, randomHex string) (bool, error) {
	passwordHash, err := m.getAuthHash(randomHex)
	if err != nil {
		return false, fmt.Errorf("failed to validate the auth answer: %v", err)
	}
	return passwordHash == resAuthAnswer.Parameters[constants.ParamAuthCLIRandAnswer], nil
}
