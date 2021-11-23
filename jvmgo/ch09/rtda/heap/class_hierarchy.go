package heap

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
//注意，粗体部分是原来的代码，其余都是新增代码。由于篇幅 限制，就不详细解释这个函数了，请读者阅读Java虚拟机规范的 8.6.5节对instanceof和checkcast指令的描述。需要注意的是:
//·数组可以强制转换成Object类型(因为数组的超类是Object)。 ·数组可以强制转换成Cloneable和Serializable类型(因为数组实
//现了这两个接口)。 ·如果下面两个条件之一成立，类型为[]SC的数组可以强制转
//换成类型为[]TC的数组:
//·TC和SC是同一个基本类型。 ·TC和SC都是引用类型，且SC可以强制转换成TC。
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

	return false
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



// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}
