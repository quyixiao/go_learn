package lang

import "unsafe"
import "go_learn/jvmgo/ch11/native"
import "go_learn/jvmgo/ch11/rtda"

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
	native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
	native.Register(jlObject, "notifyAll", "()V", notifyAll)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	//把对象引用(Object结构体指针)转换成uintptr类型，然后强制 转换成int32推入操作数栈顶。
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
//继续编辑Object.go，实现clone()方法
func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	//如果类没有实现Cloneable接口，则抛出CloneNotSupportedException异常
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	//否则调用Object结构体的 Clone()方法克隆对象，然后把对象副本引用推入操作数栈顶。
	frame.OperandStack().PushRef(this.Clone())
}

// public final native void notifyAll();
// ()V
func notifyAll(frame *rtda.Frame) {
	// todo
}
