package heap

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

// self extends c
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
// 判断S是否是T的子类，实际上也就是判断T是否是S的(直接或 间接)超类
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends iface
//判断S是否实现了T接口，就看S或S的(直接或间接)超类是否 实现了某个接口T'，T'要么是T，要么是T的子接口。
//isSubInterfaceOf()方法也在class_hierarchy.go文件中
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
