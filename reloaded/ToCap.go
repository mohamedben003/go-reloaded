package reloaded

import "strings"

func ToCap(s string) string {
	lowered := strings.ToLower(s)

	runes := []rune(lowered)

	if len(runes) > 0 {
		firstChar := strings.ToUpper(string(runes[0]))
		return firstChar + string(runes[1:])
	}
	return s
}