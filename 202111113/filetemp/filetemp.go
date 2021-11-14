package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.TempDir())      ///var/folders/t7/75ts_gyx6h3g14vdt8xgdk3m0000gn/T/
	fmt.Println(os.UserCacheDir()) ///Users/quyixiao/Library/Caches <nil>
	fmt.Println(os.UserHomeDir())  ///Users/quyixiao <nil>
	fmt.Println(os.Getwd())		//获取当前文件所在目录		/Users/quyixiao/go/src/go_learn <nil>


}
