package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p *int = new(int)
	fmt.Printf("%T\n", p)

	// ポインタ型を返す場合はnew関数、そうでない場合はmake関数
	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
func one(x *int) {
	*x = 1
}

func sample1() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)
}
