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
		for i := 1; i < len(res)-1; i++ {
			switch res[i] {
			case "(bin)":
				res[i-1] = reloaded.ConvertBase(res[i-1], res[i])
				res = reloaded.Remove(res, i)

			case "(hex)":
				res[i-1] = reloaded.ConvertBase(res[i-1], res[i])
				res = reloaded.Remove(res, i)

			case "(cap)":
				res[i-1] = reloaded.ToUpper(res[i-1])
				res = reloaded.Remove(res, i)

			case "(up)":
				res[i-1] = strings.ToUpper(res[i-1])
				res = reloaded.Remove(res, i)

			case "(low)":
				res[i-1] = strings.ToLower(res[i-1])
				res = reloaded.Remove(res, i)

			}
		}
		a = append(a, string(reloaded.Punctuations([]byte(strings.Join(res, " ")))))
	}
	// fmt.Println(a)
	joined := []byte(strings.Join(a, "\n"))

	fmt.Println(string(joined))
}

// out := ""

// err = os.WriteFile(arg[1], []byte(str), 0666)
// if err != nil {
// 	log.Fatal("err")
// }
