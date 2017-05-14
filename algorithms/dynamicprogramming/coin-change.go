package main

import (
	"fmt"
	"log"
)

type key struct {
	value, depth int64
}

var coinCounts [50]int64
var coinValues [50]int64
var nCoins int64
var n int64
var cache map[key]int64

func main() {
	fmt.Scanf("%d %d", &n, &nCoins)
	for i := int64(0); i < nCoins; i++ {
		fmt.Scan(&coinValues[i])
	}
	cache = make(map[key]int64)
	count := change(n, 0)
	fmt.Println(count)
}

func change(value int64, depth int64) (count int64) {
	if count, ok := cache[key{value, depth}]; ok {
		return count
	}
	count = true_change(value, depth)
	cache[key{value, depth}] = count
	return
}

func true_change(value int64, depth int64) (count int64) {
	currentCoins := coinValues[depth:nCoins]
	if len(currentCoins) == 0 {
		return 0
	}
	if value == 0 {
		log.Println("pim")
		return 1
	} else if value < 0 {
		return 0
	}

	// check how many ways we can change value not using this coin
	count += change(value, depth+1)
	// check how many ways we can change value using one of this coin
	count += change(value-currentCoins[0], depth)
	return
}
