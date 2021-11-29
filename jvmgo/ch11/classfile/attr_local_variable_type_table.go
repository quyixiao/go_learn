package classfile

/*
LocalVariableTypeTable_attribute {
    u2 attribute_name_index;	表示字符串"LocalVariableTypeTable"
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;		用来表示一个有效的非限定名，以指代这个局部变量
        u2 signature_index;		必须是对常量池表的一个有效索引，常量池在该索引处的成员必须是CONSTANT_Utf8_info 结构，此结构是用来表示源程序中的局部变量类型的字段签名

        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}
*/
type LocalVariableTypeTableAttribute struct {
	localVariableTypeTable []*LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

func (self *LocalVariableTypeTableAttribute) readInfo(reader *ClassReader) {
	localVariableTypeTableLength := reader.readUint16()
	self.localVariableTypeTable = make([]*LocalVariableTypeTableEntry, localVariableTypeTableLength)
	for i := range self.localVariableTypeTable {
		self.localVariableTypeTable[i] = &LocalVariableTypeTableEntry{
			startPc:        reader.readUint16(),
			length:         reader.readUint16(),
			nameIndex:      reader.readUint16(),
			signatureIndex: reader.readUint16(),
			index:          reader.readUint16(),
		}
	}
}
