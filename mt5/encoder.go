package mt5

import (
	"golang.org/x/text/encoding/unicode"
)

var utf16LEDecoder = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
var utf16LEEncoder = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()

// ToUTF16LE encodes a UTF-8 string to UTF-16 little endian encoding
func ToUTF16LE(source string) (string, error) {
	return utf16LEEncoder.String(source)
}

// ToUTF8 decodes a UTF-16 little endian string to UTF-8 encoding
func ToUTF8(source string) (string, error) {
	return utf16LEDecoder.String(source)
}
