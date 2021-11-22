package constants

import (
	"go_learn/jvmgo/ch06/instructions/base"
	"go_learn/jvmgo/ch06/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
