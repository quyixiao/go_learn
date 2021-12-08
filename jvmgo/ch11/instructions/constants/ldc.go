package constants

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"

// Push item from run-time constant pool
// ldc 和ldc_w 指令用于访问运行时常量池中的值，这包括类String的实例，但不包括double和long类型的实例
// 当运行时常量池中的条目过多时，需要使用ldc_w 指令取代ldc指令来访问常量池，ldc2_w 指令用于访问类型为double和long的运行时常量池项目
// 这个指令没有非宽版本。

type LDC struct{ base.Index8Instruction }


// Push long or double from run-time constant pool (wide index)
// 每个ldc2_w 指令的操作数都必须是常量池表内的一个有效索引，被此索引引用的常量池成员必须是CONSTANT_Long或CONSTANT_Double
// 另外，紧随其后的那个常量池索引也必须是对常量池的一个有效索引，并且该索引处的常量池成员不允许使用
type LDC2_W struct{ base.Index16Instruction }
// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string: // 如果ldc试图从运行时常量池中加载字符串常量，则先通过常量拿到Go字符串，然后把它转成Java字符串实例并把引用推入操作数栈顶。
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}


func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
