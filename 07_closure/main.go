package main

import "fmt"

func main() {
	//counter := incrementGenerator()
	//fmt.Println(counter())
	//fmt.Println(counter())
	//fmt.Println(counter())

	decremtCounter := decrementGenerator()
	fmt.Println(decremtCounter())
	fmt.Println(decremtCounter())
	fmt.Println(decremtCounter())
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func sample1() {
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func decrementGenerator() func() int {
	x := 10
	return func() int {
		x--
		return x
	}
}
