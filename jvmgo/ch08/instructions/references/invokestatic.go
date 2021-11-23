package references

import (
	"go_learn/jvmgo/ch08/instructions/base"
	"go_learn/jvmgo/ch08/rtda"
	"go_learn/jvmgo/ch08/rtda/heap"
)

// Invoke a class (static) method
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	//假定解析符号引用后得到方法M。M必须是静态方法，否则抛 出Incompatible-ClassChangeError异常
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := resolvedMethod.Class()
	//先判 断类的初始化是否已经开始，如果还没有，则需要调用类的初始化 方法，并终止指令执行。但是由于此时指令已经执行到了一半，
	//也 就是说当前帧的nextPC字段已经指向下一条指令，所以需要修改 nextPC，让它重新指向当前指令。Frame结构体的RevertNextPC()方 法做了这样的操作
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}
