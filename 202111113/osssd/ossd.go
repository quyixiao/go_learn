package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//os.Stdin 标准的输出
	fmt.Println("xxx")

	os.Stdout.Write([]byte("bbbbbbb\n")) //bbbbbbb

	/*	bytes := make([]byte, 5)
		os.Stdin.Read(bytes)
		n, err := os.Stdin.Read(bytes)
		fmt.Println(n, err, bytes)
	*/

	//可以输入空格了
	//abc oiewoiewio  oiaid
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())

	// 标准的输入输出
	fmt.Fprintf(os.Stdout,"%T",1)

}
