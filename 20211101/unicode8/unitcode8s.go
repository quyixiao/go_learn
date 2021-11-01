package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "中华人民共和国"
	fmt.Println(len(s))					//21
	fmt.Println(utf8.RuneCountInString(s))			//7
	
}
