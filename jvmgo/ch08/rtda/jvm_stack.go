package rtda

// jvm stack
type Stack struct {
	maxSize uint						//maxSize字段保存栈的容量(最多可以容纳多少帧)
	size    uint			//size字段保 存栈的当前大小，
	_top    *Frame // stack is implemented as linked list _top字段保存栈顶指针
}

//经典的链表(linked list)数据结构来实现Java虚拟机栈，这样 栈就可以按需使用内存空间，而且弹出的帧也可以及时被Go的垃 圾收集器回收
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
//push()方法把帧推入栈顶
func (self *Stack) push(frame *Frame) {
	//如果栈已经满了，按照Java虚拟机规范，应该抛出 StackOverflowError异常。在第10章才会讨论异常，
	//这里先调用 panic()函数终止程序执行
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

//如果此时栈是空的，肯定是我们的虚拟机有bug，调用panic() 函数终止程序执行即可
//top()方法只是返回栈顶帧，但并不弹出
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--

	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	return self._top
}




func (self *Stack) isEmpty() bool {
	return self._top == nil
}
