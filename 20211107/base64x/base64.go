package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	x := base64.StdEncoding.EncodeToString([]byte("i am kk"))
	fmt.Println(x)
	bytes, err := base64.StdEncoding.DecodeString(x)
	fmt.Println(string(bytes), err)
}
