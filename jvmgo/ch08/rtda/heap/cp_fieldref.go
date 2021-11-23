package heap

import "go_learn/jvmgo/ch08/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

//Field结构体比较简单，目前所有信息都是从ClassMember中继 承过来的。newFields()函数根据class文件的字段信息创建字段表
//field字段缓存解析后的字段指针，newFieldRef()方法创建 FieldRef实例
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

// jvms 5.4.3.2
//如果类D想通过字段符号引用访问类C的某个字段，首先要解 析符号引用得到类C，然后根据字段名和描述符查找字段。如果字
//段查找失败，则虚拟机抛出NoSuchFieldError异常。如果查找成功， 但D没有足够的权限访问该字段，则虚拟机抛出IllegalAccessError异 常。
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

//首先在C的字段中查找。如果找不到，在C的直接接口递归应 用这个查找过程。如果还找不到的话，在C的超类中递归应用这个 查找过程。如果仍然找不到，则查找失败。
// 用通俗的语言描述字段访问规则。如果字段是public，则任何 类都可以访问。如果字段是protected，则只有子类和同一个包下的 类可以访问。如果字段有默认访问权限
//(非public，非protected，也 非privated)，则只有同一个包下的类可以访问。否则，字段是 private，只有声明这个字段的类才能访问。
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
