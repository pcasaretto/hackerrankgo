package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e6 + 6

func main() {
	var size int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size = toInt(scanner.Bytes())
	A := make([]int, size)
	B := make([]int, size)
	for i := 0; i < size; i++ {
		scanner.Scan()
		A[i] = toInt(scanner.Bytes())
	}
	for i := 0; i < size; i++ {
		scanner.Scan()
		B[i] = toInt(scanner.Bytes())
	}
	countA := make([]int, N)
	lmulA := make([]int, N)
	countB := make([]int, N)
	lmulB := make([]int, N)
	for i := 0; i < size; i++ {
		countA[A[i]] += 1
	}
	for i := 1; i <= N; i++ {
		for j := i; j < N; j += i {
			if countA[j] != 0 {
				lmulA[i] = j
			}
		}
	}
	for i := 0; i < size; i++ {
		countB[B[i]] += 1
	}
	for i := 1; i < N; i++ {
		for j := i; j < N; j += i {
			if countB[j] != 0 {
				lmulB[i] = j
			}
		}
	}
	mx := 0
	for i := 1; i < N; i++ {
		if lmulA[i] > 0 && lmulB[i] > 0 {
			mx = i
		}
	}
	fmt.Println(lmulA[mx] + lmulB[mx])
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
