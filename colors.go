package check

import (
	"os"
	"strings"
)

var colorsStart = "\033[31m"
var colorsEnd = "\033[39m"

func init() {
	if showColors() {
		return
	}
	noColors()
}

func noColors() {
	colorsStart = ""
	colorsEnd = ""
}

func showColors() bool {
	var useColors = true

	if os.Getenv("NO_COLOR") != "" {
		useColors = false
	}
	if strings.Contains(os.Getenv("TERM"), "dumb") {
		useColors = false
	}
	return useColors
}
