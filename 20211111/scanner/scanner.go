package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("user.log")

	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			//99382322222222222
			//9832983
			//89i32832
			//dsiodia
			//iodsoia
		}
	}
}
