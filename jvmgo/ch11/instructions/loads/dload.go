package loads

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"

// Load double from local variable
// 加载指令从局部变量表获取变量，然后推入操作数栈顶
type DLOAD struct{ base.Index8Instruction }

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, self.Index)
}

type DLOAD_0 struct{ base.NoOperandsInstruction }

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperandsInstruction }

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperandsInstruction }

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperandsInstruction }

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
