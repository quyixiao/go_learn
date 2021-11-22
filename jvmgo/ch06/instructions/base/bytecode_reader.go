package base

type BytecodeReader struct {
	code []byte // bytecodes
	pc   int
}

//code字段存放字节码，pc字段记录读取到了哪个字节。为了避 免每次解码指令都新创建一个BytecodeReader实例，
//给它定义一个 Reset()方法
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

//下面实现一系列的Read()方法。先看最简单的ReadUint8()方 法
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}


//ReadInt16()方法调用ReadUint16()，然后把读取到的值转成 int16返回
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}
//ReadUint16()连续读取两字节
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

//ReadInt32()方法连续读取4字节
func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// used by lookupswitch and tableswitch
//还需要定义两个方法:ReadInt32s()和SkipPadding()。这两个 方法只有tableswitch和lookupswitch指令使用
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// used by lookupswitch and tableswitch
//tableswitch指令操作码的后面有0~3字节的padding，以保证 defaultOffset在字节码中的地址是4的倍数
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
