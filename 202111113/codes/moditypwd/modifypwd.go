package main

import (
	"crypto/md5"
	"encoding/csv"
	"fmt"
	"os"
)


func getMd5(str string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}

const passwordfile = "pwdfile.csv"

func modifyPassword(originpwd, newpwd string) {
	if file, err :=  os.Open(passwordfile); err == nil {
		defer file.Close()
		reader := csv.NewReader(file)
		line, _ := reader.Read()
		if getMd5(originpwd) == line[0] {
			password(newpwd)
		}
	}
}
func password(newpwd string ){
	if fileNew, err := os.Create(passwordfile); err == nil {
		writer := csv.NewWriter(fileNew)
		defer fileNew.Close()
		md5Pwd := getMd5(newpwd);
		fmt.Println(md5Pwd)
		writer.Write([]string{md5Pwd})
		writer.Flush()
	}
}



func main() {

	fmt.Println(getMd5("123456"))

	modifyPassword("123456","aa")
}
