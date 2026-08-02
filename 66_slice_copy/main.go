package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Println(dst, "コピー数:", copied)

	// dstを変更してもsrcには影響しない
	dst[0] = 999
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)
}
