package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var current int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	current = toInt(scanner.Bytes())
	current++
	for !test(current) {
		current++
	}
	fmt.Println(current)
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func test(ticket int) bool {
	left := ticket%1e6/1e5 +
		ticket%1e5/1e4 +
		ticket%1e4/1e3
	right := ticket%1e3/1e2 +
		ticket%1e2/1e1 +
		ticket%1e1/1e0
	return left == right
}
