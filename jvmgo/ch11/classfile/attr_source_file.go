package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;		项引用的字符串表示被编译的class文件的源文件的名字，
不要把它理解成源文件所在目录或源文件的绝对路径名，这些与平台相关的附加信息，必须由运行时解释器（runtime interpreter）或
开发工具的实际使用文件名时提供
}
*/
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
