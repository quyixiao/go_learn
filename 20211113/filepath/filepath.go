package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Abs(".")) ///Users/quyixiao/go/src/go_learn <nil>
	fmt.Println(os.Args[0])        ///private/var/folders/t7/75ts_gyx6h3g14vdt8xgdk3m0000gn/T/GoLand/___go_build_go_learn_202111113_filepath
	//获取当前运行程序的绝对路径
	fmt.Println(filepath.Abs(os.Args[0])) ///private/var/folders/t7/75ts_gyx6h3g14vdt8xgdk3m0000gn/T/GoLand/___go_build_go_learn_202111113_filepath <nil>
	fmt.Println(filepath.Base("test"))    //获取文件名
	path, _ := filepath.Abs(os.Args[0])
	fmt.Println(filepath.Base(path)) //___go_build_go_learn_202111113_filepath
	//获取父目录的地址
	fmt.Println(filepath.Dir(path)) ///private/var/folders/t7/75ts_gyx6h3g14vdt8xgdk3m0000gn/T/GoLand
	fmt.Println(filepath.Ext(path)) //获取文件的扩展名或文件后缀

	dirpath := filepath.Dir(path)

	initpath := dirpath + "config/ip.init"
	fmt.Println(initpath)

	fmt.Printf("%T ,%#v \n", initpath, initpath)      //string ,"/private/var/folders/t7/75ts_gyx6h3g14vdt8xgdk3m0000gn/T/GoLandconfig/ip.init"
	fmt.Println(filepath.Join(dirpath, "abc", "dev")) ///private/var/folders/t7/75ts_gyx6h3g14vdt8xgdk3m0000gn/T/GoLand/abc/dev

	// glob 是用来找文件的
	fmt.Println(filepath.Glob("./*/*/*/*.java")) //[20211113/filepath/xxx/bb.java] <nil>

	//遍历一个目录下的所有文件
	filepath.Walk(".", func(path string, fileinfo os.FileInfo, error error) error {
		//	fmt.Println(path, fileinfo.Name())
		return nil
	})

}
