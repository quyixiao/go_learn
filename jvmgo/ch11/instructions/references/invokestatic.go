package references

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"

// Invoke a class (static) method
// invokestatic 指令用于调用命名类中的类方法 （static方法）
// invokedynamic指令用于调用以绑定invokedynamic指令的调用点对象(call site object)作为目标的方法，调用点对象是一个特殊的语法
//结构，当一条invokedynamic指令首次被java虚拟机执行前，Java虚拟机将会执行一个引导方法（bootstrapmethod）并以这个方法的执行结果
// 作为调用点对象，因此，每条invokedynamic指令都有独一无二的链接状态，这是它与其他方法调用指令的一个差异
// invokestatic指令调用静态方法
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	//	通过这个索引， 可以从当前类的运行时常量池中找到一个方法符号引用，解析这个 符号引用就可以得到一个方法。
	resolvedMethod := methodRef.ResolvedMethod()

	//假定解析符号引用后得到方法M。M必须是静态方法，否则抛 出Incompatible-ClassChangeError异常。M不能是类初始化方法。

	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()
	//类 初始化方法只能由Java虚拟机调用，不能使用invokestatic指令调 用。这一规则由class文件验证器保证，这里不做检查。
	//如果声明M 的类还没有被初始化，则要先初始化该类。
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	//对于invokestatic指令，M就是最终要执行的方法，调用 InvokeMethod()函数执行该方法。
	base.InvokeMethod(frame, resolvedMethod)
}
