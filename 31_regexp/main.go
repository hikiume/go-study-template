package main

import (
	"fmt"
	"regexp"
)

func main() {
	// URLによって処理を変えたい時に使用できる
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}

func sampleMatchString() {
	match, _ := regexp.MatchString("a([a-z0-9]+)e", "app0e")
	fmt.Println(match)
}

// 同じ正規表現を繰り返す場合はMustCompileの関数のほうが高速
func sampleMustCompile() {
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)
}

// FindString : 正規表現に一致する最も左側の文字列を返す
// URLの後半のような文字列が正規表現に一致するかを確認できる
func sampleFindString() {
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
}
