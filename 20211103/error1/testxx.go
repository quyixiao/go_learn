package main

import (
	"fmt"
	"strings"
)

func main() {
	inputId := "zhangsan 1 "
	fmt.Scan(&inputId)
	xx := strings.Split(inputId, "_")
	fmt.Println(xx)
}
