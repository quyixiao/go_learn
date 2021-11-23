package stack

import (
	"go_learn/jvmgo/ch09/instructions/base"
	"go_learn/jvmgo/ch09/rtda"
)

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
pop指令把栈顶变量弹出
*/
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
pop指令只能用于弹出int、float等占用一个操作数栈位置的变 量。double和long变量在操作数栈中占据两个位置，需要使用pop2 指令弹出
*/
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
