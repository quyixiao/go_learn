package constants

import (
	"go_learn/jvmgo/ch07/instructions/base"
	"go_learn/jvmgo/ch07/rtda"
)

// Push byte
type BIPUSH struct {
	val int8
}

//以bipush指令为例，FetchOperands()和Execute()方法的代码
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// Push short
type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
