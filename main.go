package main

import (
	"fmt"
	"os"
	"strings"
	"reloaded/reloaded"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}

	r, err := os.ReadFile(string(arg[0]))
	if err != nil {
		fmt.Println(err)
		return
	}

	res1 := reloaded.TrimmedSlice(r)
	a := []string{}

	for j := 0; j < len(res1); j++ {
		res := res1[j]
		
		res = reloaded.Core(res)
		// fmt.Println(res)
		
		a = append(a, string(reloaded.Punctuations([]byte(strings.Join(res, " ")))))
	}

	joined := []byte(strings.Join(a, "\n"))
	fmt.Println(string(joined))
}