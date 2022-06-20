package mt5

import (
	"fmt"
	"strings"
)

type MT5Command struct {
	Command         string
	Encrypted       bool
	Parameters      map[string]interface{}
	Body            string
	ResponseHasBody bool
}

func (c *MT5Command) toString() string {
	i := 0
	components := make([]string, len(c.Parameters))
	for k, v := range c.Parameters {
		val := fmt.Sprintf("%v", v)
		components[i] = fmt.Sprintf("%s=%v", k, c.sanitizeValue(val))
		i++
	}
	return c.Command + "|" + strings.Join(components, "|") + "\r\n" + c.Body
}

func (c *MT5Command) sanitizeValue(value string) string {
	value = strings.Replace(value, "\\", "\\\\", -1)
	value = strings.Replace(value, "=", "\\=", -1)
	value = strings.Replace(value, "|", "\\|", -1)
	return strings.Replace(value, "\n", "\\\n", -1)
}
