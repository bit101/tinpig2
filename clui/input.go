// Package clui has command line ui functions
package clui

import (
	"bufio"
	"os"
	"strings"

	"github.com/bit101/go-ansi"
)

// ReadStringDefault displays a prompt and collects input.
func ReadStringDefault(prompt, def string) string {
	ansi.Printf(ansi.Yellow, "%s ", prompt)
	if def != "" {
		ansi.Printf(ansi.Blue, "(%s) ", def)
	}
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	if str == "" {
		return def
	}
	return strings.TrimSuffix(str, "\n")
}

// ReadString displays a prompt and collects input.
func ReadString(prompt string) string {
	ansi.Printf(ansi.Yellow, "%s ", prompt)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSuffix(str, "\n")
}