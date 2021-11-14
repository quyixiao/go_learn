package main

import (
	"encoding/gob"
	"os"
)

func main() {
	users := map[int]string{1: "张三", 2: "李四"}
	file, err := os.Create("user.gob")
	if err == nil {
		defer file.Close()
		encoder := gob.NewEncoder(file)
		encoder.Encode(users)
	}
}
