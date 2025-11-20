package reloaded

import (
	"fmt"
	"strings"
)

func TrimmedSlice(r []byte) [][]string {
	str := strings.Split(string(r), "\n")
	res := [][]string{}

	for i := 0; i < len(str); i++ {
		x := strings.Fields(str[i])
		x = an(x)
		res = append(res, x)
	}
	fmt.Println(res)
	return res
}
