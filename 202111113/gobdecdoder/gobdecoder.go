package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {

	users := map[int]string{}

	file, err := os.Open("user.gob")
	if err == nil {
		defer file.Close()
		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)
		fmt.Println(users)				//map[1:张三 2:李四]


	}

}
