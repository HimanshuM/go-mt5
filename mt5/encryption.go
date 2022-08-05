package mt5

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/HimanshuM/go-mt5/constants"
)

// parseHexString parses a hex string into string built from decimal byte array
func parseHexString(srvRand string) (string, error) {
	srvRandByteArr := make([]byte, 0)
	srvRandRune := []rune(srvRand)
	for i := 0; i < len(srvRandRune); i += 2 {
		hexRune := srvRandRune[i : i+2]
		hexStr := string(hexRune)
		decimal, err := strconv.ParseInt(hexStr, 16, 32)
		if err != nil {
			return "", fmt.Errorf("failed to parse %s: %v", constants.PARAM_AUTH_SRV_RAND, err)
		}
		srvRandByteArr = append(srvRandByteArr, byte(decimal))
	}
	return string(srvRandByteArr), nil
}

// getRandomHex generates a hex string of specified length from random decimal numbers
func getRandomHex(size int) string {
	hexString := ""
	for i := 0; i < size; i++ {
		hexString += fmt.Sprintf("%02x", rand.Intn(254))
	}
	return hexString
}
