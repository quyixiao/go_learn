package loads

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"

// Load reference from local variable
// 加载指令从局部变量表获取变量，然后推入操作数栈顶。
// aload系列指令 操作引用类型变量
type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
