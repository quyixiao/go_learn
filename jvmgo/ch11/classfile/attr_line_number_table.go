package classfile

/*
LineNumberTable_attribute {
    u2 attribute_name_index;		用来表示字符串"LineNumberTable"
    u4 attribute_length;
    u2 line_number_table_length;	数组的每个成员都表明源文件中的行号会在code数组中的哪一条指令处发生变化，line_number_table每个成员都具有如下两项
    {   u2 start_pc;			项的值必须是code[]数组的一个索引，code[]在该索引的指令码，表示源文件中新的行的起点
        u2 line_number;			项的值必须与源文件中对应的行号相匹配
    } line_number_table[line_number_table_length];
}
*/

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func (self *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(self.lineNumberTable) - 1; i >= 0; i-- {
		entry := self.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
