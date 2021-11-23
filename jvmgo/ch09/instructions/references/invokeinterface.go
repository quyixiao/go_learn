package references

import "go_learn/jvmgo/ch09/instructions/base"
import "go_learn/jvmgo/ch09/rtda"
import "go_learn/jvmgo/ch09/rtda/heap"

// Invoke interface method
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	//先拿到当前类、当前常量池、方法符号引用，然后解析符号引 用，拿到解析后的类和方法。
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	//假定从方法符号引用中解析出来的类是C，方法是M。如果M 是构造函数，则声明M的类必须是C，否则抛出NoSuchMethodError
	//异常。如果M是静态方法，则抛出IncompatibleClassChangeError异 常。
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//从操作数栈中弹出this引用，如果该引用是null，抛出 NullPointerException异常。注意，在传递参数之前，不能破坏操作 数栈的状态。
	//给OperandStack结构体添加一个GetRefFromTop()方 法，该方法返回距离操作数栈顶n个单元格的引用变量。
	//比如 GetRefFromTop(0)返回操作数栈顶引用，GetRefFromTop(1)返回从 栈顶开始的倒数第二个引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException") // todo
	}

	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}


	//上面这段代码比较难懂，把它翻译成更容易理解的语言:如果 调用的中超类中的函数，但不是构造函数，且当前类的 ACC_SUPER标志被设置，
	//需要一个额外的过程查找最终要调用的 方法;否则前面从方法符号引用中解析出来的方法就是要调用的方 法。
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	//如果查找过程失败，或者找到的方法是抽象的，抛出 AbstractMethodError异常。
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	//上面的判断确保protected方法只能被声明该方法的类或子类 调用。如果违反这一规定，则抛出IllegalAccessError异常
	//最后，如果一切正常，就调用方法。这里 之所以这么复杂，是因为调用超类的(非构造函数)方法需要特别处 理。
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
