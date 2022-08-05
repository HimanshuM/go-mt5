package mt5

import (
	"fmt"
	"strings"
)

// Command struct defines a command that can be sent to the MT5 server
type Command struct {
	Command         string
	Encrypted       bool
	Parameters      map[string]interface{}
	Body            string
	ResponseHasBody bool
}

// toString strigfyies the MT5Command into raw command string that would be sent the MT5 server
func (c *Command) toString() string {
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
func (c *Command) sanitizeValue(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "=", "\\=")
	value = strings.ReplaceAll(value, "|", "\\|")
	return strings.ReplaceAll(value, "\n", "\\\n")
}
