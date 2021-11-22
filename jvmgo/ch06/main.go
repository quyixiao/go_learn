package main

import (
	"fmt"
	"go_learn/jvmgo/ch06/classpath"
	"go_learn/jvmgo/ch06/rtda/heap"
	"strings"
)

// #  ch05 -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_131.jdk/Contents/Home/jre"    GaussTest
//pc: 7 inst:*comparisons.IF_ICMPGT &{{13}}
//LocalVars:[{0 <nil>} {5050 <nil>} {101 <nil>}]
//OperandStack:&{0 [{101 <nil>} {100 <nil>}]}
//panic: Unsupported opcode: 0xb2! [recovered]
//        panic: Unsupported opcode: 0xb2!
//
//goroutine 1 [running]:
//main.catchErr(0xc0000103c0)
//        /Users/quyixiao/go/src/go_learn/jvmgo/ch06/interpreter.go:34 +0xf9
func main() {
	//cmd := parseCmd()

	cmd := &Cmd{
		helpFlag:    false,
		versionFlag: false,
		cpOption:    "/Users/quyixiao/go/src/go_learn/jvmgo/ch06",
		class:       "MyObject",
		XjreOption:  "/Library/Java/JavaVirtualMachines/jdk1.8.0_131.jdk/Contents/Home/jre",
		args:        []string{},
	}
	fmt.Printf("%#v \n", cmd)
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		fmt.Println("------------")
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
