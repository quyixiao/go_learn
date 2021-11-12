package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func copyfile(src, dest string) {
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close()
		destfile, err := os.Create(dest)
		if err != nil {
			fmt.Println(err)
		} else {
			defer destfile.Close()
			bytes := make([]byte, 1024*1024*10)

			reader := bufio.NewReader(srcfile)
			writer := bufio.NewWriter(destfile)

			for {
				n, err := reader.Read(bytes)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break

				}
				writer.Write(bytes[:n])
				writer.Flush()
			}
		}
	}
}

func main() {
	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "dest file")
	help := flag.Bool("h", false, "bool")
	flag.Usage = func() {
		fmt.Println(`Usage：copyfile -s -scrfile -d.destfile 
Options : 
`)
		flag.PrintDefaults()
	}
	flag.Parse()
	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else {
		copyfile(*src, *dest)
	}
	fmt.Println("%T ,%T ")
}
