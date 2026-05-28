package main

import (
	"fmt"
	"sync"
)

func normal(s string) {
	for range 5 {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	// バッファあり
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}

// バッファなし
func unbufferdChannel() {

	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i)
	}
}
