package rtda

import "go_learn/jvmgo/ch09/rtda/heap"

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
	stack *Stack
	// todo
}


//和堆一样，Java虚拟机规范对Java虚拟机栈的约束也相当宽 松。
//Java虚拟机栈可以是连续的空间，也可以不连续;可以是固定大 小，也可以在运行时动态扩展 [1] 。
//如果Java虚拟机栈有大小限制， 且执行线程所需的栈空间超出了这个限制，会导致 StackOverflowError异常抛出。
//如果Java虚拟机栈可以动态扩展，但 是内存已经耗尽，会导致OutOfMemoryError异常抛出。
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

//PushFrame()和PopFrame()方法只是调用Stack结构体的相应 方法而已
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

//CurrentFrame()方法返回当前帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}



func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
