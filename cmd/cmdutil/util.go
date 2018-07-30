package cmdutil

import (
	"os"
	"strings"
)

// LastMatchArgOptionName return last option name that arg matched.
func LastMatchArgOptionName(a string) (opt string) {
	for i, arg := range os.Args {
		if arg == a {
			opt = strings.TrimLeft(os.Args[i-1], "-")
		}
	}

	return opt
}
