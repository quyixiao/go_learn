package base

import "go_learn/jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)		//FetchOperands()方法从字节码中提取操作数
	Execute(frame *rtda.Frame)						//Execute()方法 执行指令逻辑
}


//NoOperandsInstruction表示没有操作数的指令，所以没有定义
type NoOperandsInstruction struct {
	// empty
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

//BranchInstruction表示跳转指令，Offset字段存放跳转偏移量。
type BranchInstruction struct {
	Offset int
}

//FetchOperands()方法从字节码中读取一个uint16整数，转成int后赋 给Offset字段
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

//继续编辑instruction.go文件，在其中定义Index8Instruction结构 体
type Index8Instruction struct {
	Index uint
}

//存储和加载类指令需要根据索引存取局部变量表，索引由单 字节操作数给出。
//把这类指令抽象成Index8Instruction结构体，用 Index字段表示局部变量表索引。
//FetchOperands()方法从字节码中 读取一个int8整数，转成uint后赋给Index字段
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}


//有一些指令需要访问运行时常量池，常量池索引由两字节操 作数给出。把这类指令抽象成Index16Instruction结构体，
//用Index字 段表示常量池索引。FetchOperands()方法从字节码中读取一个 uint16整数，转成uint后赋给Index字段
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
