package comparisons

import (
	"go_learn/jvmgo/ch08/instructions/base"
	"go_learn/jvmgo/ch08/rtda"
)

// Compare float
type FCMPG struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag { //也就是说，当两个float变量中至少有一个是NaN时，用fcmpg指 令比较的结果是1，而用fcmpl指令比较的结果是-1。
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
