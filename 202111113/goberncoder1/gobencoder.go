package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
}

func main() {
	users := []User{
		User{1, "张三", time.Now(), "184899832", "上海"},
		User{2, "李四", time.Now(), "184899832", "上海"},
		User{3, "王五", time.Now(), "184899832", "上海"},
		User{4, "赵六", time.Now(), "184899832", "上海"},
	}
	file, err := os.Create("user1.gob")
	if err == nil {
		defer file.Close()
		encoder := gob.NewEncoder(file)
		encoder.Encode(users)
	}

	file1, err := os.Open("user1.gob")

	users1 := []User{}
	if err == nil {
		defer file1.Close()
		decoder := gob.NewDecoder(file1)
		decoder.Decode(&users1)
		fmt.Println(users1) //map[1:{1 张三 2021-11-13 17:57:31.116331 +0800 CST 184899832 上海} 2:{2 李四 2021-11-13 17:57:31.116334 +0800 CST 184899832 上海} 3:{3 王五 2021-11-13 17:57:31.116334 +0800 CST 184899832 上海} 4:{4 赵17:57:31.116334 +0800 CST 184899832 上海}]
	}

}
