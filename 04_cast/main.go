package main

import (
	"fmt"
	"strconv"
)

// string型からint型への変換
func main() {
	var s = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v \n", i, i)
}
