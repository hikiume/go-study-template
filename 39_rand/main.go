package main

import (
	"math/rand"
)

func main() {

	for range 10 {
		// 0 ~ 9 までの乱数が生成される
		randomNumber := rand.Intn(10)
		println(randomNumber)
	}
}
