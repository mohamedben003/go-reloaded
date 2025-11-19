package reloaded

import "strings"

func TrimmedSlice(r []byte) [][]string {
	str := strings.Split(string(r), "\n")
	res := [][]string{}

	for i := 0; i < len(str); i++ {
		x := strings.Split(str[i], " ")
		for j :=0 ; j<len(x); j++ {
			if x[j] == "" {
				x = append(x[:j], x[j+1:]...)
				j--
			}
		}
		res = append(res, x)
	}
	return res
}
