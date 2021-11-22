package references

import "go_learn/jvmgo/ch07/instructions/base"
import "go_learn/jvmgo/ch07/rtda"
import "go_learn/jvmgo/ch07/rtda/heap"

// Fetch field from object
type GET_FIELD struct{ base.Index16Instruction }

//getfield指令获取对象的实例变量值，然后推入操作数栈，它需 要两个操作数。第一个操作数是uint16索引，用法和前面三个指令 一样。
//第二个操作数是对象引用，用法和putfield一样。
func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	//弹出对象引用，如果是null，则抛出NullPointerException
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

	//根据字段类型，获取相应的实例变量值，然后推入操作数栈。 至此，getfield指令也解释完毕了
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}
