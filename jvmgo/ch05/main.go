package main

import (
	"fmt"
	"go_learn/jvmgo/ch02/classpath"
	"go_learn/jvmgo/ch05/classfile"
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
//        /Users/quyixiao/go/src/go_learn/jvmgo/ch05/interpreter.go:34 +0xf9
func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
