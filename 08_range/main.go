package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	for _, v := range m {
		fmt.Println(v)
	}
}

func sample() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
}

func sampleRange() {
	l := []string{"python", "go", "java"}

	for i, v := range l {
		fmt.Println(i, v)
	}
}

func sampleRange2() {
	l := []string{"python", "go", "java"}

	for _, v := range l {
		fmt.Println(v)
	}
}

func sampleMap() {
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
