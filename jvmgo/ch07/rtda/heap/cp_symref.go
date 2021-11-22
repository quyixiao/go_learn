package heap

// symbolic reference
type SymRef struct {
	cp        *ConstantPool						//cp字段存放符号引用所在的运行时常量池指针，这样就可以通 过符号引用访问到运行时常量池，进一步又可以访问到类数据。
	className string					//className字段存放类的完全限定名。
	class     *Class				//class字段缓存解析后的类结 构体指针，这样类符号引用只需要解析一次就可以了，后续可以直 接使用缓存值。
	//对于类符号引用，只要有类名，就可以解析符号引 用。对于字段，首先要解析类符号引用得到类数据，然后用字段名 和描述符查找字段数据。方法符号引用的解析过程和字段符号引用 类似。
}

//如果类符号引用已经解析，ResolvedClass()方法直接返回类指 针，否则调用resolveClassRef()方法进行解析。
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// jvms8 5.4.3.1
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	//通俗地讲，如果类D通过符号引用N引用类C的话，要解析N， 先用D的类加载器加载C，然后检查D是否有权限访问C，如果没
	//有，则抛出IllegalAccessError异常。Java虚拟机规范5.4.4节给出了类 的访问控制规则，把这个规则翻译成Class结构体的 isAccessibleTo()方法
	//也就是说，如果类D想访问类C，需要满足两个条件之一:C是 public，或者C和D在同一个运行时包内。
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
