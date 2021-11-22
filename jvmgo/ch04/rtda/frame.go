package rtda

// stack frame
type Frame struct {
	lower        *Frame        //lower字段用来实现链表数据结构
	localVars    LocalVars     //localVars字段保存 局部变量表指针
	operandStack *OperandStack //operandStack字段保存操作数栈指针。
	// todo
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
