package references

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
// invokespecial指令用于调用一些需要特殊处理的实例方法，包括实现初始化方法（见2.9节）私有方法和父类方法
//  invokespecial指令用来调 用无须动态绑定的实例方法，包括构造函数、私有方法和通过super 关键字调用的超类方法。
// invokespecial指 令也比较好理解。首先，因为私有方法和构造函数不需要动态绑 定，所以invokespecial指令可以加快方法调用速度。
// 其次，使用super 关键字调用超类中的方法不能使用invokevirtual指令，否则会陷入 无限循环。
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	//先拿到当前类、当前常量池、方法符号引用，然后解析符号引用，拿到解析后的类和方法
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	//假定从方法符号引用中解析出来的类是C，方法是M。如果M 是构造函数，则声明M的类必须是C，否则抛出NoSuchMethodError异常。
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	//如果M是静态方法，则抛出IncompatibleClassChangeError异常。
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//从操作数栈中弹出this引用，如果该引用是null，抛出 NullPointerException异常。注意，在传递参数之前，不能破坏操作 数栈的状态。
	//给OperandStack结构体添加一个GetRefFromTop()方 法，该方法返回距离操作数栈顶n个单元格的引用变量。比如
	//GetRefFromTop(0)返回操作数栈顶引用，GetRefFromTop(1)返回从 栈顶开始的倒数第二个引用，
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	//判断确保protected方法只能被声明该方法的类或子类调用。如果违反这一规定，则抛出IllegalAccessError异常
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod
	//如果调用的中超类中的函数，但不是构造函数，且当前类的 ACC_SUPER标志被设置，需要一个额外的过程查找最终要调用的
	//方法;否则前面从方法符号引用中解析出来的方法就是要调用的方法。
	/**
	public class AccSuperDemo {

	    // lib1.jar
	    public static class A {
	        public void foo() {
	            System.out.println("A.foo");
	        }
	    }
	    public static class B extends A {
	        // empty
	    }

	    // lib2.jar
	    public static class C extends B {
	        public void foo() {
	            super.foo();
	        }
	    }

	    public static void main(String[] args) {
	        new C().foo();
	    }

	}
	console : A.foo
	*/
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {

		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}

	// 如果查找过程失败，或者找到的方法是抽象的，抛出 AbstractMethodError异常。
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
