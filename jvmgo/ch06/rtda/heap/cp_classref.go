package heap

import "go_learn/jvmgo/ch06/classfile"

type ClassRef struct {
	SymRef
}

//ClassRef继承了SymRef，但是并没有添加任何字段。
//newClassRef()函数根据class文件中存储的类常量创建ClassRef实 例
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
