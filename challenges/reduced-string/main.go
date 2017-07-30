package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	res := reduce(string(b))
	if len(res) > 0 {
		fmt.Println(res)
		return
	}
	fmt.Println("Empty String")
}

func reduce(s string) string {
	// log.Println("reduce", s)
	l := len(s)
	if l == 1 {
		return s
	}
	left := reduce(s[:l/2])
	right := reduce(s[l/2:])
	return join(left, right)
}

func join(left, right string) string {
	// log.Println("join", left, right)
	leftLen := len(left)
	if leftLen > 0 && len(right) > 0 && left[leftLen-1] == right[0] {
		return join(left[:leftLen-1], right[1:])
	}
	return left + right
}
