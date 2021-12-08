package references

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"

// Determine if object is of given type
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	////第一个操作数是uint16索引， 从方法的字节码中获取，通过这个索引可以从当前类的运行时常量 池中找到一个类符号引用。
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	//第二个操作数是对象引用，从操作数栈 中弹出。
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	//先弹出对象引用，如果是null，则把0推入操作数栈。用Java代 码解释就是，如果引用obj是null的话，不管ClassYYY是哪种类型， 下面这条if判断都是false
	//if (obj instanceof ClassYYY) {...}
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
