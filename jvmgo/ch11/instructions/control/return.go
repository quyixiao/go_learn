package control

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"
// 当返回的值是boolean ，byte ,char ,short 和int类型的时候使用
type IRETURN struct{ base.NoOperandsInstruction }
// 返回的是一个double类型
type DRETURN struct{ base.NoOperandsInstruction }
// 返回的是一个符点类型
type FRETURN struct{ base.NoOperandsInstruction }
// 返回的是一个long类型
type LRETURN struct{ base.NoOperandsInstruction }
// return指令供声明为void的方法，实例初始化方法，类和接口的初始化方法使用
type RETURN struct{ base.NoOperandsInstruction }
// Return reference from method
type ARETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}



func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}



func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}


func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
