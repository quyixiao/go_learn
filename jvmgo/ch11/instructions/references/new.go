package references

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"

// Create new object
type NEW struct{ base.Index16Instruction }

//new指令的操作数是一个uint16索引，来自字节码。通过这个索 引，可以从当前类的运行时常量池中找到一个类符号引用。
//解析这个类符号引用，拿到类数据，然后创建对象，并把对象引用推入栈顶，new指令的工作就完成了。
func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted() {
		frame.RevertNextPC()
		//另外，如果解析后的类还没有初始化，则需要先初始化类。 在第7章实现方法调用之后会详细讨论类的初始化，这里暂时先忽 略。
		base.InitClass(frame.Thread(), class)
		return
	}
	//因为接口和抽象类都不能实例化，所以如果解析后的类是接 口或抽象类，按照Java虚拟机规范规定，需要抛出InstantiationError 异常。
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
