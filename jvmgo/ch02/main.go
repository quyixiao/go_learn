package main

import (
	"fmt"
	"go_learn/jvmgo/ch02/classpath"
	"strings"
)



//startJVM 先打印出命令行参数，然后读取主类数据，并打印到控制台，虽然还是无法真正启动Java虚拟机，不过相比第1章，已经有很大的
//进步，打开命令行窗口，执行下面的命令编译本章代码
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath :%s class:%s args :%v \n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s \n", cmd.class)
		return
	}
	fmt.Printf("class data:%v \n", classData)
}


func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1 ")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}

}
