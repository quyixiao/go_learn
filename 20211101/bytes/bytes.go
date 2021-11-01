package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte{'a','b','c'}
	fmt.Printf("%T %#v\n", b) // []uint8 %!v(MISSING)
	fmt.Println(b)            // [97 98 99]


	s := string(b)
	fmt.Printf("%T %#v \n",s,s) // string "abc"


	fmt.Println(bytes.Compare([]byte{'a','b'},[]byte{'c','d'})) // -1

	fmt.Println(bytes.Index([]byte("abcdef"),[]byte("def"))) //3
	fmt.Println(bytes.Contains([]byte("abcdef"),[]byte("defd")))	 //false

}
