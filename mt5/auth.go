package mt5

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
)

func (m *MT5) Auth() error {
	resAuthStart, err := m.sendAuthStart()
	if err != nil {
		return err
	}
	srvRand, found := resAuthStart.Parameters[PARAM_SRV_RAND]
	if !found {
		return fmt.Errorf("response param %s not found in response", PARAM_SRV_RAND)
	}
	passwordHash, err := m.getAuthHash(srvRand.(string))
	if err != nil {
		return err
	}
	if _, err = m.sendAuthAnswer(passwordHash); err != nil {
		return err
	}
	m.connected = true
	return nil
}

func (m *MT5) sendAuthStart() (*MT5Response, error) {
	cmd := &MT5Command{
		Command: CMD_AUTH_START,
		Parameters: map[string]interface{}{
			"VERSION":      API_VERSION,
			"AGENT":        WORD_API,
			"LOGIN":        m.config.Username,
			"TYPE":         WORD_MANAGER,
			"CRYPT_METHOD": CRYPT_METHOD,
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

func (m *MT5) sendAuthAnswer(passwordHash string) (*MT5Response, error) {
	cmd := &MT5Command{
		Command: CMD_AUTH_ANSWER,
		Parameters: map[string]interface{}{
			PARAM_SRV_RAND_ANSWER: passwordHash,
			PARAM_CLI_RAND:        getRandomHex(16),
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

func (m *MT5) getAuthHash(srvRand string) (string, error) {
	utf16LEPassword, err := ToUTF16LE(m.config.Password)
	if err != nil {
		return "", err
	}
	passwordHash := md5.Sum([]byte(utf16LEPassword))
	saltedPassword := string(passwordHash[:]) + WORD_API
	saltedPasswordHash := md5.Sum([]byte(saltedPassword))
	srvRandStr, err := getSrvRandByteArray(srvRand)
	if err != nil {
		return "", err
	}
	finalString := string(saltedPasswordHash[:]) + srvRandStr
	finalHash := md5.Sum([]byte(finalString))
	finalHashHex := ""
	for _, each := range finalHash {
		finalHashHex += fmt.Sprintf("%02x", each)
	}
	return finalHashHex, nil
}

func getSrvRandByteArray(srvRand string) (string, error) {
	srvRandByteArr := make([]byte, 0)
	srvRandRune := []rune(srvRand)
	for i := 0; i < len(srvRandRune); i += 2 {
		hexRune := srvRandRune[i : i+2]
		hexStr := string(hexRune)
		decimal, err := strconv.ParseInt(hexStr, 16, 32)
		if err != nil {
			return "", fmt.Errorf("failed to parse %s: %v", PARAM_SRV_RAND, err)
		}
		srvRandByteArr = append(srvRandByteArr, byte(decimal))
	}
	return string(srvRandByteArr), nil
}

func getRandomHex(len int) string {
	hexString := ""
	for i := 0; i < len; i++ {
		hexString += fmt.Sprintf("%02x", rand.Intn(254))
	}
	return hexString
}
