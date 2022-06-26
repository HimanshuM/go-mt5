package mt5

import (
	"fmt"
	"strings"
)

// MT5Command struct defines a command that can be sent to the MT5 server
type MT5Command struct {
	Command         string
	Encrypted       bool
	Parameters      map[string]interface{}
	Body            string
	ResponseHasBody bool
}

// toString strigfyies the MT5Command into raw command string that would be sent the MT5 server
func (c *MT5Command) toString() string {
	i := 0
	components := make([]string, len(c.Parameters))
	for k, v := range c.Parameters {
		val := fmt.Sprintf("%v", v)
		components[i] = fmt.Sprintf("%s=%v", k, c.sanitizeValue(val))
		i++
	}
	return fmt.Sprintf("%s|%s\r\n%s", c.Command, strings.Join(components, "|"), c.Body)
}

// sanitizeValue sanitizes the command parameter values
func (c *MT5Command) sanitizeValue(value string) string {
	value = strings.Replace(value, "\\", "\\\\", -1)
	value = strings.Replace(value, "=", "\\=", -1)
	value = strings.Replace(value, "|", "\\|", -1)
	return strings.Replace(value, "\n", "\\\n", -1)
}
