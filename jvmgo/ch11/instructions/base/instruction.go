package base

import "go_learn/jvmgo/ch11/rtda"

type Instruction interface {
	//Execute()方法 执行指令逻辑。有很多指令的操作数都是类似的。为了避免重复代 码，按照操作数类型定义一些结构体，并实现FetchOperands()方 法。
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
	// empty
}

// FetchOperands()方法从字节码中提取操作数
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

type BranchInstruction struct {
	Offset int
}

//BranchInstruction表示跳转指令，Offset字段存放跳转偏移量。 FetchOperands()方法从字节码中读取一个uint16整数，转成int后赋 给Offset字段
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
