package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileName := fmt.Sprintf("%s.log", time.Now().Format("2006-1-2 15-04-05"))
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	logger := log.New(logFile, "", log.Llongfile)
	logger.Println("Test logs")
}
