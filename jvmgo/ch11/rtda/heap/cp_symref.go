package heap

// symbolic reference
type SymRef struct {
	//cp字段存放符号引用所在的运行时常量池指针，这样就可以通 过符号引用访问到运行时常量池，进一步又可以访问到类数据。
	cp        *ConstantPool
	className string					//className字段存放类的完全限定名。
	class     *Class		//class字段缓存解析后的类结 构体指针，这样类符号引用只需要解析一次就可以了，后续可以直 接使用缓存值。
}

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
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
