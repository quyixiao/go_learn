package extended

import (
	"go_learn/jvmgo/ch05/instructions/base"
	"go_learn/jvmgo/ch05/rtda"
)

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

//goto_w指令和goto指令的唯一区别就是索引从2字节变成了4 字节。
func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
