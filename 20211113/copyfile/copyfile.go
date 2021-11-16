package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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
func copyDir(src, dest string) {
	files, err := ioutil.ReadDir(src)
	fmt.Println("==============",err)
	if err == nil {

		os.MkdirAll(dest, os.ModePerm)
		for _, file := range files {
			if file.IsDir() {
				copyDir(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			} else {
				fmt.Println(file.Name())
				copyfile(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			}
		}
	}

}

func main() {


	copyDir("/Users/quyixiao/Desktop/subline","/Users/quyixiao/Desktop/testx")
}
