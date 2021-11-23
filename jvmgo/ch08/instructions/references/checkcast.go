package references

import "go_learn/jvmgo/ch08/instructions/base"
import "go_learn/jvmgo/ch08/rtda"
import "go_learn/jvmgo/ch08/rtda/heap"

// Check whether object is of given type
type CHECK_CAST struct{ base.Index16Instruction }


//如果对象引用不是null，则解析类符号引用，判断对象是否是 类的实例，然后把判断结果推入操作数栈。Java虚拟机规范给出了 具体的判断步骤，
//我们在Object结构体的IsInstanceOf()方法中实 现，稍后给出代码。
//instanceof指令 会改变操作数栈(弹出对象引用，推入判断结果);checkcast则不改 变操作数栈(如果判断失败，直接抛出ClassCastException异常)。
func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	//先从操作数栈中弹出对象引用，再推回去，这样就不会改变操 作数栈的状态。如果引用是null，则指令执行结束。也就是说，null 引用可以转换成任何类型
	if ref == nil {
		return
	}
	//则解析类符号引用，判断对象是否是 类的实例。如果是的话，指令执行结束，
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	//否则抛出 ClassCastException。
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
