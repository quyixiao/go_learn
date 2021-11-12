package main

import (
	"bufio"
	"os"
)

func main() {
	file ,err := os.Create("user.log")
	if err == nil {
		defer file.Close()
		write := bufio.NewWriter(file)
		write.WriteString("abcidsoidsoi\n")
		write.WriteString("wriiodsoiiods")
		write.Flush()
	}
}
