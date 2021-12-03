package rtda

import "go_learn/jvmgo/ch11/rtda/heap"

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/
type Thread struct {
	pc    int // the address of the instruction currently being executed
	stack *Stack			//stack字段 是Stack结构体(Java虚拟机栈)指针
	// todo
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
func (self *Thread) ClearStack() {
	//它调用了Stack结构体的clear()方法
	self.stack.clear()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}
