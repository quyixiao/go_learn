package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func md5Method(reader *bufio.Reader) string {
	bytes := make([]byte, 1024*1024*10)
	hasher := md5.New()
	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		hasher.Write(bytes[:n])
	}
	return fmt.Sprintf("%x", hasher.Sum([]byte(nil)))
}

func md5file(path string) string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	} else {
		return md5Method(bufio.NewReader(file))
	}
	return ""
}
func mid5str(txt string) string {
	//return fmt.Sprintf("%x", md5.Sum([]byte(txt)))
	return md5Method(bufio.NewReader(strings.NewReader(txt)))

}
func main() {
	txt := flag.String("s", "", "md5 txt")
	path := flag.String("f", "", "md5 txt")
	help := flag.Bool("h", false, " help")
	flag.Usage = func() {
		fmt.Println("Usage : md5.ext [-s 123abc] -f filepath")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help || *txt == "" && *path == "" {
		flag.Usage()
	} else {
		var md5 string
		if *path != "" {
			md5 = md5file(*path)
		} else {
			md5 = mid5str(*txt)
		}
		fmt.Println(md5)
	}

}
