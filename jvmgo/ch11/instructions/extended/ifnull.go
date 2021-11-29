package extended

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }
// Branch if reference not null          ifnull和ifnonnull指令把栈顶的 引用弹出。
type IFNONNULL struct{ base.BranchInstruction }

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
