package main

import (
	"fmt"
	"go_learn/jvmgo/ch05/classfile"
	"go_learn/jvmgo/ch05/instructions"
	"go_learn/jvmgo/ch05/instructions/base"
	"go_learn/jvmgo/ch05/rtda"
)


//interpret()方法的其余代码先创建一个Thread实例，然后创建 一个帧并把它推入Java虚拟机栈顶，最后执行方法。
func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()		//得到Code属性之后，可以进一步获得执行方法所需的局部变 量表和操作数栈空间
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}


//回到interpret()方法，我们的解释器目前还没有办法优雅地结 束运行。因为每个方法的最后一条指令都是某个return指令，而还 没有实现return指令，
//所以方法在执行过程中必定会出现错误，此 时解释器逻辑会转到catchErr()函数
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}


//把局部变量表和操作数栈的内容打印出来，以此来观察方法 的执行结果。还剩一个loop()函数
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)					//loop()函数循环执行“计算pc、解码指令、执行指令”这三个步 骤，直到遇到错误!代码中有一个函数还没有给出代码: NewInstruction()
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
