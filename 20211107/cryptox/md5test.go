package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	x := md5.Sum([]byte("i amkk"))
	y := fmt.Sprintf("%X \n", x)
	fmt.Println(y)
	//
	z := hex.EncodeToString(x[:])
	fmt.Println(z)

	m := md5.New()
	m.Write([]byte("i"))
	m.Write([]byte(" amkk"))
	fmt.Printf("%x\n", m.Sum(nil))

}
