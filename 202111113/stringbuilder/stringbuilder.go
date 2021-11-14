package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("abc"))
	builder.Write([]byte("def"))
	fmt.Println(builder.String()) //abcdef
}
