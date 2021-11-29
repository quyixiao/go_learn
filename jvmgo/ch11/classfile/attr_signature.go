package classfile

/*
Signature_attribute {
    u2 attribute_name_index;				必需是常量池表中的一个有效索引
    u4 attribute_length;					attribute_length项的值必需为2
    u2 signature_index;						用以表示类签名，方法类型签名或字段类型签名，如果signatrue属性是
	ClassFile结构的属性，则这个结构表示类签名，如果当前Signatrue属性是method_info结构的属性，则这个结构表示方法类型签名
	如果当前的Signatrue属性是field_info结构的属性，则这个结构是表示字段类型签名
	签名（Signatrue）用来编码以Java语言所写的声明，这些声明使用了Java虚拟机类型系统之外的类型，在只能访问class文件的情况下
签名有助于实现反射，调试及编译
	如果类，接口，构造器，方法或字段的声明使用了类变量或参数化类型，那么Java编译器就必需为此生成签名，具体来说，Java编译器
在这些情况下必需生成签名。
	1.当类声明或接口声明具备泛型形式，或者其超类或超接口是参数化类型，又或者前两条兼备时，必须生成类签名
	2.当方法声明或构造器声明具备了泛型形式，或者其形式参数类型或返回类型是类型变量或参数化类型，或者在throws子句中使用了类型变量
	又或者前述三条具备两条或者三者兼备时，必需生成方法签名

	如果方法声明或构造器声明了throws子句不涉及类型变量，那么编译器在生成方法签名时，可以将该声明为不含throws子句的声明
	当字段声明，形式参数声明或局部变量声明中的类型使用了类型参数或参数化类型时，必需为该声明字段签名。
	引用类型签名。

}
*/
type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (self *SignatureAttribute) readInfo(reader *ClassReader) {
	self.signatureIndex = reader.readUint16()
}

func (self *SignatureAttribute) Signature() string {
	return self.cp.getUtf8(self.signatureIndex)
}
