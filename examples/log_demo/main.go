package main

import (
	"log"
	"os"
)

var (
	Error   *log.Logger
	Warning *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// log.Println("message")
	// // log.Fatalln 会调用os.Exit(1)强制退出程序
	// log.Fatalln("fatal message")
	// log.Panicln("panic message")

	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
