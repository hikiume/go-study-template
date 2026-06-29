package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) (string, error) {
	if age < 0 {
		return "", errors.New("年齢に負の数は指定できません")
	}
	if age < 18 {
		return "", errors.New("18歳未満はアクセスできません")
	}
	return "アクセス許可", nil
}

func main() {
	result, err := checkAge(17)
	if err != nil {
		fmt.Println("【エラー発生】:", err)
		return
	}

	fmt.Println(result)
}
