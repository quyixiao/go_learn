package heap

import "go_learn/jvmgo/ch11/classfile"

type ExceptionTable []*ExceptionHandler

//ExceptionTable只是[]*ExceptionHandler的别名而已
type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ClassRef
}

//newExceptionTable()函数把class文件中的异常处理表转换成 ExceptionTable类型。有一点需要特别说明:异常处理项的catchType 有可能是0。
//我们知道0是无效的常量池索引，但是在这里0并非表 示catch-none，而是表示catch-all，它的用法马上就会看到。
//getCatchType()函数从运行时常量池中查找类符号引用
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}

	return table
}

//getCatchType()函数从运行时常量池中查找类符号引用
func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil // catch all
	}
	return cp.GetConstant(index).(*ClassRef)
}

//异常处理表查找逻辑前面已经描述过，此处不再赘述。这里注 意两点。第一，startPc给出的是try{}语句块的第一条指令，
//endPc给 出的则是try{}语句块的下一条指令。第二，如果catchType是nil(在 class文件中是0)，表示可以处理所有异常，这是用来实现finally子句的。
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}
