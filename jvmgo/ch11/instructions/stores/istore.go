package stores

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"

// Store int into local variable
//和加载指令刚好相反，存储指令把变量从操作数栈顶弹出，然 后存入局部变量表
type ISTORE struct{ base.Index8Instruction }

func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, uint(self.Index))
}

type ISTORE_0 struct{ base.NoOperandsInstruction }

func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct{ base.NoOperandsInstruction }

func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct{ base.NoOperandsInstruction }

func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct{ base.NoOperandsInstruction }

func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
// istore 命令的作用是从操作栈中弹出值，并存储到本地变量表中
func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
