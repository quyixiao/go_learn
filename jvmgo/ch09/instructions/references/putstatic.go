package references

import "go_learn/jvmgo/ch09/instructions/base"
import "go_learn/jvmgo/ch09/rtda"
import "go_learn/jvmgo/ch09/rtda/heap"

// Set static field in class
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	//先拿到当前方法、当前类和当前常量池，然后解析字段符号引
	//用。如果声明字段的类还没有被初始化，则需要先初始化该类
	class := field.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}


	//如果解析后的字段是实例字段而非静态字段，则抛出 IncompatibleClassChangeError异常。
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//如果是final字段，则实际操作的 是静态常量，只能在类初始化方法中给它赋值。否则，会抛出 IllegalAccessError异常。类初始化方法由编译器生成，名字是 <clinit>
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	//根据字段类型从操作数栈中弹出相应的值，然后赋给静态变量。
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// todo
	}
}
