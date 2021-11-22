package control

import (
	"go_learn/jvmgo/ch07/instructions/base"
	"go_learn/jvmgo/ch07/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
