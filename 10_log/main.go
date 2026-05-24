package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	LoggingSettings("test.log")
	_, err := os.Open("fdafdsafa")

	if err != nil {
		log.Println("エラーが発生しました:", err)
	}
}

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func sampleLog() {
	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	log.Fatal("error!!")

	fmt.Println("ok!")
}

func sampleError() {
	_, err := os.Open("fdafdsafa")
	if err != nil {
		// エラーがある場合にプログラムを終了させる
		log.Fatalln("Exit", err)
	}
	log.Println("logging!")

	fmt.Println(err)
}
