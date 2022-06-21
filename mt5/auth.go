package mt5

import (
	"crypto/md5"
	"fmt"
)

func (m *MT5) Auth() error {
	if m.connected {
		return nil
	}

	if !m.config.isCryptMethodKnown() {
		return fmt.Errorf("unknown encryption method: %s", m.config.CryptMethod)
	}
	resAuthStart, err := m.sendAuthStart()
	if err != nil {
		return err
	}
	srvRand, found := resAuthStart.Parameters[PARAM_AUTH_SRV_RAND]
	if !found {
		return fmt.Errorf("response param %s not found in response", PARAM_AUTH_SRV_RAND)
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
		if randomCrypt, present := resAuthAnswer.Parameters[PARAM_AUTH_CRYPT_RAND]; !present {
			return fmt.Errorf("auth answer response does not contain %s", PARAM_AUTH_CRYPT_RAND)
		} else {
			m.randomCrypt = randomCrypt.(string)
		}
	}
	m.connected = validResponse
	return nil
}

func (c *MT5Config) isCryptMethodKnown() bool {
	return c.CryptMethod == CRYPT_METHOD_DEFAULT || c.CryptMethod == CRYPT_METHOD_NONE
}

func (m *MT5) sendAuthStart() (*MT5Response, error) {
	cmd := &MT5Command{
		Command: CMD_AUTH_START,
		Parameters: map[string]interface{}{
			"VERSION":      API_VERSION,
			"AGENT":        WORD_API,
			"LOGIN":        m.config.Username,
			"TYPE":         WORD_MANAGER,
			"CRYPT_METHOD": m.config.CryptMethod,
		},
	}
	response, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, fmt.Errorf("Auth failed at %s: %v", CMD_AUTH_START, err)
	}
	if response.CommandName != CMD_AUTH_START {
		return nil, fmt.Errorf("response of %s (%d) is invalid: %s (%d)", CMD_AUTH_START, len(CMD_AUTH_START), response.CommandName, len(response.CommandName))
	}
	if response.ReturnCode != 0 {
		return nil, fmt.Errorf("authorization failed: %v", response.ReturnValue)
	}
	return response, nil
}

func (m *MT5) sendAuthAnswer(passwordHash string, randomHex string) (*MT5Response, error) {
	cmd := &MT5Command{
		Command: CMD_AUTH_ANSWER,
		Parameters: map[string]interface{}{
			PARAM_AUTH_SRV_RAND_ANSWER: passwordHash,
			PARAM_AUTH_CLI_RAND:        randomHex,
		},
	}
	response, err := m.IssueCommand(cmd)
	if err != nil {
		return nil, fmt.Errorf("auth failed at %s: %v", CMD_AUTH_ANSWER, err)
	}
	if response.CommandName != CMD_AUTH_ANSWER {
		return nil, fmt.Errorf("response of %s (%d) is invalid: %s (%d)", CMD_AUTH_ANSWER, len(CMD_AUTH_ANSWER), response.CommandName, len(response.CommandName))
	}
	if response.ReturnCode != 0 {
		return nil, fmt.Errorf("authorization failed: %v", response.ReturnValue)
	}
	return response, nil
}

func (m *MT5) getAuthHash(hexString string) (string, error) {
	utf16LEPassword, err := ToUTF16LE(m.config.Password)
	if err != nil {
		return "", err
	}
	passwordHash := md5.Sum([]byte(utf16LEPassword))
	saltedPassword := string(passwordHash[:]) + WORD_API
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

func (m *MT5) validateAuthAnswer(resAuthAnswer *MT5Response, randomHex string) (bool, error) {
	passwordHash, err := m.getAuthHash(randomHex)
	if err != nil {
		return false, fmt.Errorf("failed to validate the auth answer: %v", err)
	}
	return passwordHash == resAuthAnswer.Parameters[PARAM_AUTH_CLI_RAND_ANSWER], nil
}
