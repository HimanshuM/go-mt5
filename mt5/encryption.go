package mt5

import (
	"fmt"
	"math/rand"
	"strconv"
)

func parseHexString(srvRand string) (string, error) {
	srvRandByteArr := make([]byte, 0)
	srvRandRune := []rune(srvRand)
	for i := 0; i < len(srvRandRune); i += 2 {
		hexRune := srvRandRune[i : i+2]
		hexStr := string(hexRune)
		decimal, err := strconv.ParseInt(hexStr, 16, 32)
		if err != nil {
			return "", fmt.Errorf("failed to parse %s: %v", PARAM_AUTH_SRV_RAND, err)
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
