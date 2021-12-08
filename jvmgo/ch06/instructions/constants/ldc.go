package constants

import (
	"go_learn/jvmgo/ch06/instructions/base"
	"go_learn/jvmgo/ch06/rtda"
)

// Push item from run-time constant pool
type LDC struct{ base.Index8Instruction }
// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }
// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }


func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}


func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}



//先从当前类的运行时常量池中取出常量。如果是int或float常 量，则提取出常量值，则推入操作数栈。其他情况还无法处理，暂时调用panic()
//函数终止程序执行。ldc_2w指令的Execute()方法单独实现
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
