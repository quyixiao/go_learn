package main

import (
	"fmt"
	"strings"
)
import "go_learn/jvmgo/ch11/instructions"
import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"

func interpret(thread *rtda.Thread, logInst bool) {
	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

//修改之后interpret()函数简单了许多，直接调用loop()函数进 入循环即可。loop()函数循环执行“计算pc、解码指令、执行指令”这三个步 骤，直到遇到错误!
//1. 在每次循环开始，先拿到当前帧
//2. pc从当前方法中解码出一条指令
//3. 指令执行完毕之后
//4. 判断Java虚拟机栈中是否还有 帧。如果没有则退出循环;

// 4条指令都修改完毕了，但是新增加的代码做了些什么?先判断类的初始化是否已经开始，如果还没有，则需要调用类的初始化方法，
// 并终止指令执行。但是由于此时指令已经执行到了一半，也就是说当前帧的nextPC字段已经指向下一条指令，所以需要修改 nextPC，
// 让它重新指向当前指令。Frame结构体的RevertNextPC()方法做了这样的操作
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {

		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())



		classNameB := frame.Method().Class().Name()

		if(strings.Contains(classNameB,"jvmgo/book/ch06/MyObject")){
				fmt.Println("==========MyObject========")
		}

		logInstruction(frame, inst)


		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		lineNum := method.GetLineNumber(frame.NextPC())
		fmt.Printf(">> line:%4d pc:%4d %v.%v%v \n",
			lineNum, frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
