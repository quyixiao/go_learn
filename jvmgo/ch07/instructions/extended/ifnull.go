package extended

import (
	"go_learn/jvmgo/ch07/instructions/base"
	"go_learn/jvmgo/ch07/rtda"
)

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

//根据引用是否是null进行跳转，ifnull和ifnonnull指令把栈顶的 引用弹出。
func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}


func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
