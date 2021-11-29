package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;				用以表示字符串"Synthetic"
    u4 attribute_length;					attribute_length固定为0
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
