package references

import "go_learn/jvmgo/ch09/instructions/base"
import "go_learn/jvmgo/ch09/rtda"

// Get length of array
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	//如果数组引用是null，则需要抛出NullPointerException异常，否 则取数组长度，推入操作数栈顶即可。下面实现<t>aload和 <t>astore系列指令。
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
