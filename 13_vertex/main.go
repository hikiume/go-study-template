package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v := &Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)
}

// ポインタを使用することで値を変更することができる
func changeVertex(v *Vertex) {
	v.X = 1000
}

func sample2() {
	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

}

func sample1() {
	v := Vertex{X: 1}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v.X = 100
	fmt.Println(v.X, v.Y)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)
}
