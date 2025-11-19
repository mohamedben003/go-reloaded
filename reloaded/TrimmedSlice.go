package reloaded

import "strings"

func TrimmedSlice(r []byte) []string {
	str := strings.Split(string(r), " ")
	res := []string{}
	for i := 0; i < len(str); i++ {
		if str[i] != "" {
			res = append(res, str[i])
		}
	}
	return res
}
