package references

import "go_learn/jvmgo/ch09/instructions/base"
import "go_learn/jvmgo/ch09/rtda"

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
// invokespecial指令及invokestatic指令的索引字节rwkt，必须表示指向常量池表的有效索引，如果class文件的版本号小于52,那么
// 该索引引用的常量池项目，其类型必须是CONSTANT_Methodref，如果class文件版本号大于等于52，那该索引所引用的常量池项，
// 其类型必须是CONSTANT_Methodref 或CONSTANT_InterfaceMethodref
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
