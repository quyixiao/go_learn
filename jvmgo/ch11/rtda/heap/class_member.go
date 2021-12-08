package heap

import "go_learn/jvmgo/ch11/classfile"

type ClassMember struct {
	accessFlags    uint16
	name           string
	descriptor     string
	signature      string
	annotationData []byte // RuntimeVisibleAnnotations_attribute
	class          *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Signature() string {
	return self.signature
}
func (self *ClassMember) AnnotationData() []byte {
	return self.annotationData
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// jvms 5.4.4
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	//如果字段是public，则任何 类都可以访问。
	if self.IsPublic() {
		return true
	}
	c := self.class
	//如果字段是protected，则只有子类和同一个包下的 类可以访问。
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	//如果字段有默认访问权限(非public，非protected，也 非privated)，则只有同一个包下的类可以访问。否则，字段是 private，只有声明这个字段的类才能访问。
	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}
