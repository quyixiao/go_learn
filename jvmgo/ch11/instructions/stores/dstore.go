package stores

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"

// Store double into local variable
//和加载指令刚好相反，存储指令把变量从操作数栈顶弹出，然 后存入局部变量表
type DSTORE struct{ base.Index8Instruction }

func (self *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(self.Index))
}

type DSTORE_0 struct{ base.NoOperandsInstruction }

func (self *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOperandsInstruction }

func (self *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOperandsInstruction }

func (self *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOperandsInstruction }

func (self *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
