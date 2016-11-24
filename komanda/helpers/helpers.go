package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

// Contains item in array of strings
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// FormatMessage adds support for bold text
func FormatMessage(str string) string {
	r, err := regexp.Compile("\\*[a-zA-Z]+\\*")

	if err != nil {
		return str
	}

	m := r.FindAllString(str, -1)

	for _, v := range m {
		str = strings.Replace(str, v,
			fmt.Sprintf(
				"\x1b[38;5;%d;%dm%s\x1b[0m",
				15, 1, v,
			),
			1,
		)
	}

	return str
}
