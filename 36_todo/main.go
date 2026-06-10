package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	var todoList []Task
	currentID := 1

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("--- Go Todo アプリへようこそ ---")

	for {
		fmt.Println("\n操作を選んでください:")
		fmt.Println("1: タスクを追加")
		fmt.Println("2: タスクを一覧表示")
		fmt.Println("3: タスクを完了にする")
		fmt.Println("4: 終了")
		fmt.Println("選択 > ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			fmt.Print("タスク名を入力してください: ")
			if scanner.Scan() {
				title := strings.TrimSpace(scanner.Text())
				if title == "" {
					fmt.Println("エラー: タスク名が空です。")
					continue
				}

				newTask := Task{ID: currentID, Title: title, Done: false}
				todoList = append(todoList, newTask)
				fmt.Printf("[%s]を追加しました！\n", title)
				currentID++
			}
		case "2":
			fmt.Println("タスク一覧")
			if len(todoList) == 0 {
				fmt.Println("タスクはありません")
				continue
			}

			for _, task := range todoList {
				status := " "
				if task.Done {
					status = "X"
				}
				fmt.Printf("[%s] ID:%d - %s\n", status, task.ID, task.Title)
			}
		case "3":
			fmt.Println("完了にするタスクのIDを入力してください: ")
			if scanner.Scan() {
				idStr := strings.TrimSpace(scanner.Text())
				id, err := strconv.Atoi(idStr)
				if err != nil {
					fmt.Println("エラー: 有効な数値を入力してください。")
					continue
				}

				found := false
				for i := range todoList {
					if todoList[i].ID == id {
						todoList[i].Done = true
						fmt.Printf("タスク「%s」を完了にしました！\n", todoList[i].Title)
						found = true
						break
					}
				}

				if !found {
					fmt.Println("エラー: 指定されたIDのタスクが見つかりません。")
				}
			}
		case "4":
			fmt.Println("アプリを終了します。お疲れ様でした！")
			return

		default:
			fmt.Println("1~4の数値を入力してください。")
		}
	}
}
