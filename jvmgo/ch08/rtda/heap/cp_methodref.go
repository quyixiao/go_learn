package heap

import "go_learn/jvmgo/ch08/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}



// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	//如果类D想通过方法符号引用访问类C的某个方法，先要解析
	// 符号引用得到类C。如果C是接口，则抛出 IncompatibleClassChangeError异常
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//如果找不到对应的方法，则抛出NoSuchMethodError异常，否 则检查类D是否有权限访问该方法。如果没有，则抛出 IllegalAccessError异常。
	//isAccessibleTo()方法是在ClassMember结构 体中定义的，在第6章就已经实现了。
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}




func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
