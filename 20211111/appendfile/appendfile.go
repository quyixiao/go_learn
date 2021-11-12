package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	path := "/Users/quyixiao/Desktop/test/test.log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)
	fmt.Println(file, err)

	file.WriteString("aiodsids")
	file.WriteString("aiodsids")
	file.WriteString("aiodsids")
	file.WriteString("aiodsids")

	log.SetOutput(file)
	log.SetPrefix("qyx")
	log.Print(time.Now().Unix())
	file.Close()

}
