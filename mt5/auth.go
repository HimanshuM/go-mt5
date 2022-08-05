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
	srvRand, found := resAuthStart.Parameters[constants.PARAM_AUTH_SRV_RAND]
	if !found {
		return fmt.Errorf("response param %s not found in response", constants.PARAM_AUTH_SRV_RAND)
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
		randomCrypt, present := resAuthAnswer.Parameters[constants.PARAM_AUTH_CRYPT_RAND]
		if !present {
			return fmt.Errorf("auth answer response does not contain %s", constants.PARAM_AUTH_CRYPT_RAND)
		}
		m.randomCrypt = randomCrypt.(string)
	}
	m.connected = validResponse
	return nil
}

// isEncryptMethodKnown checks if the encryption method is implemented
func (c *Config) isEncryptMethodKnown() bool {
	return c.CryptMethod == constants.CRYPT_METHOD_DEFAULT || c.CryptMethod == constants.CRYPT_METHOD_NONE
}

// sendAuthStart sends the AUTH_START command to the MT5 server
func (m *Client) sendAuthStart() (*Response, error) {
	cmd := &Command{
		Command: constants.CMD_AUTH_START,
		Parameters: map[string]interface{}{
			"VERSION":      constants.API_VERSION,
			"AGENT":        constants.WORD_API,
			"LOGIN":        m.config.Username,
			"TYPE":         constants.WORD_MANAGER,
			"CRYPT_METHOD": m.config.CryptMethod,
		},
	}
	response, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, fmt.Errorf("Auth failed at %s: %v", constants.CMD_AUTH_START, err)
	}
	if response.CommandName != constants.CMD_AUTH_START {
		return nil, fmt.Errorf("response of %s (%d) is invalid: %s (%d)", constants.CMD_AUTH_START, len(constants.CMD_AUTH_START), response.CommandName, len(response.CommandName))
	}
	if response.ReturnCode != 0 {
		return nil, fmt.Errorf("authorization failed: %v", response.ReturnValue)
	}
	return response, nil
}

// sendAuthAnswer sends AUTH_ANSWER command to the MT5 server
func (m *Client) sendAuthAnswer(passwordHash, randomHex string) (*Response, error) {
	cmd := &Command{
		Command: constants.CMD_AUTH_ANSWER,
		Parameters: map[string]interface{}{
			constants.PARAM_AUTH_SRV_RAND_ANSWER: passwordHash,
			constants.PARAM_AUTH_CLI_RAND:        randomHex,
		},
	}
	response, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, fmt.Errorf("auth failed at %s: %v", constants.CMD_AUTH_ANSWER, err)
	}
	if response.CommandName != constants.CMD_AUTH_ANSWER {
		return nil, fmt.Errorf("response of %s (%d) is invalid: %s (%d)", constants.CMD_AUTH_ANSWER, len(constants.CMD_AUTH_ANSWER), response.CommandName, len(response.CommandName))
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
	saltedPassword := string(passwordHash[:]) + constants.WORD_API
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
	return passwordHash == resAuthAnswer.Parameters[constants.PARAM_AUTH_CLI_RAND_ANSWER], nil
}
