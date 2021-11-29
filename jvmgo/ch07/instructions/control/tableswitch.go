package control

import (
	"go_learn/jvmgo/ch07/instructions/base"
	"go_learn/jvmgo/ch07/rtda"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
// 编译器使用tableswith 和lookupswith 指令来生成switch语句的编译代码，tableswitch指令用于表示switch结构中的case语句块，
// 它可以高效的从索引表中确定case语句块的分支
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

/***
int chooseNear(int i) {
switch (i) {
	case 0: return 0;
	case 1: return 1;
	case 2: return 2;
	default: return -1;
} }

 */

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()							//defaultOffset对应默认情况下执行跳转所需的字节码偏移量;
	self.low = reader.ReadInt32()							//low和high记录case的取值范围;
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1			//jumpOffsets是一个索引表，里面存 放high-low+1个int值，对应各种case情况下，执行跳转所需的字节 码偏移量。
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}


//Execute()方法先从操作数栈中弹出一个int变量，然后看它是 否在low和high给定的范围之内。如果在，则从jumpOffsets表中查出 偏移量进行跳转，否则按照defaultOffset跳转
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	base.Branch(frame, offset)
}
