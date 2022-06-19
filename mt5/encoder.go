package mt5

import (
	"golang.org/x/text/encoding/unicode"
)

var utf16LEDecoder = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
var utf16LEEncoder = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()

func ToUTF16LE(source string) (string, error) {
	return utf16LEEncoder.String(source)
}

func ToUTF8(source string) (string, error) {
	return utf16LEDecoder.String(source)
}
