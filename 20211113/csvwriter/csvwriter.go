package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("user.csv")

	if err == nil {
		defer file.Close()

		writer := csv.NewWriter(file)
		writer.Write([]string{"编号", "姓名", "性别"})
		writer.Write([]string{"1", "张三", "男"})
		writer.Write([]string{"2", "李四", "男"})

		writer.Flush()

	}

	file1, err1 := os.Open("user.csv")

	if err1 == nil {
		reader := csv.NewReader(file1)
		for {
			line, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			fmt.Println(line)

		}

	}

}
