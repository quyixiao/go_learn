package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	err := ioutil.WriteFile("user.log", []byte("99382322222222222"), os.ModePerm)
	fmt.Println(err)

}
