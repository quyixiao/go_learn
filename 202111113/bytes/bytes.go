package main

import (
	"bytes"
	"fmt"
)

func main() {

	buffer := bytes.NewBuffer([]byte("abcdef"))
	buffer.Write([]byte("1234576"))
	buffer.WriteString("abcdef!")
	fmt.Println(buffer.String()) //abcdef1234576abcdef!

	bytes := make([]byte, 2)
	buffer.Read(bytes)
	fmt.Println(string(bytes)) //ab

	buffer.Read(bytes)
	fmt.Println(string(bytes)) //cd
	buffer.Read(bytes)
	fmt.Println(string(bytes)) //ef

	//line, _ := buffer.ReadBytes('!')
	line, _ := buffer.ReadString('!') //1234576abcdef!
	fmt.Println(string(line)) //1234576abcdef!

}
