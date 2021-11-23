package constants

import (
	"go_learn/jvmgo/ch09/instructions/base"
	"go_learn/jvmgo/ch09/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
