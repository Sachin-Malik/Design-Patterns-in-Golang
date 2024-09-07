package main

import (
	"fmt"
	"time"
)

type LibFunc func(val int) bool

func isEven(val int) bool {
	time.Sleep(time.Millisecond * 1000)
	return val%2 == 0
}

func Hello1() {
	fmt.Println(isEven(12))
	fmt.Println(isEven(13))
}

func cacheWrapper(innerFn LibFunc, cache *map[int]bool) LibFunc {
	return func(val int) bool {
		if val, ok := (*cache)[val]; ok {
			fmt.Println(val)
			return val
		}
		result := innerFn(val)
		(*cache)[val] = result
		fmt.Println(val)

		return result
	}
}
