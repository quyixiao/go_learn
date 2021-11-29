package references

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"

// Invoke a class (static) method
// invokestatic 指令用于调用命名类中的类方法 （static方法）
// invokedynamic指令用于调用以绑定invokedynamic指令的调用点对象(call site object)作为目标的方法，调用点对象是一个特殊的语法
//结构，当一条invokedynamic指令首次被java虚拟机执行前，Java虚拟机将会执行一个引导方法（bootstrapmethod）并以这个方法的执行结果
// 作为调用点对象，因此，每条invokedynamic指令都有独一无二的链接状态，这是它与其他方法调用指令的一个差异
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}
