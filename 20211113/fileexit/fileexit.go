package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("xxx")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在 ")
		}
	} else {
		file.Close()
	}

	fileinfo, err := os.Stat("test") // 获取文件的信息
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在 ")
		}
	} else {
		fmt.Println(fileinfo.Name())    //文件名
		fmt.Println(fileinfo.IsDir())   //是否是一个目录
		fmt.Println(fileinfo.Size())    //文件大小
		fmt.Println(fileinfo.ModTime()) //2021-11-13 15:22:33.370362518 +0800 CST 文件的修改时间

		if fileinfo.IsDir() {
			dirfile, err := os.Open("test")
			if err == nil {
				defer dirfile.Close()
				childrens, _ := dirfile.Readdir(-1) //读取所有的名字
				for _, children := range childrens {
					fmt.Println(children.Name(), children.IsDir())
					names, _ := dirfile.Readdirnames(-1)
					for _, name := range names {
						fmt.Println("-----------",name) //
					}
				}
			}
		}
	}

}
