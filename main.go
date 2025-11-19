package main

import (
	"fmt"
	"os"
	"reloaded/reloaded"
	"strings"
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
	r = reloaded.Punctuations(r)
	res := reloaded.TrimmedSlice(r)

	fmt.Println(res)
	for i := 1; i < len(res)-1; i++ {
		switch res[i] {
		case "(bin)":
			res[i-1] = reloaded.ConvertBase(res[i-1], res[i])
			res = reloaded.Remove(res, i)
			break
		case "(hex)":
			res[i-1] = reloaded.ConvertBase(res[i-1], res[i])
			res = reloaded.Remove(res, i)
			break
		case "(cap)":
			res[i-1] = reloaded.ToUpper(res[i-1])
			res = reloaded.Remove(res, i)
			break
		case "(up)":
			res[i-1] = strings.ToUpper(res[i-1])
			res = reloaded.Remove(res, i)
			break
		case "(low)":
			res[i-1] = strings.ToLower(res[i-1])
			res = reloaded.Remove(res, i)
			break
		}

	}

}

// out := ""

// err = os.WriteFile(arg[1], []byte(str), 0666)
// if err != nil {
// 	log.Fatal("err")
// }
