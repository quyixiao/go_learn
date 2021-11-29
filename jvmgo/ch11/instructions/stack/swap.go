package stack

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"

// Swap the top two operand stack values
//swap指令交换栈顶的两 个变量。
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
