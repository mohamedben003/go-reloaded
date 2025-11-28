package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reloaded/reloaded"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println("you should put both the sample and the result file names")
		return
	}
	
	if filepath.Base(arg[1]) == "main.go" {
		fmt.Println("Not allowed to change the main from there 🙃")
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
		a = append(a, string(reloaded.Punctuations([]byte(strings.Join(res, " ")))))
	}

	joined := []byte(strings.Join(a, "\n"))

	err = os.WriteFile(arg[1], joined, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully to", arg[1])
}
