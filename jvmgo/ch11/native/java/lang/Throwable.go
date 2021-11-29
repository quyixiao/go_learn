package lang

import "fmt"
import "go_learn/jvmgo/ch11/native"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"

const jlThrowable = "java/lang/Throwable"

//StackTraceElement结构体用来记录Java虚拟机栈帧信息
type StackTraceElement struct {
	fileName   string								//fileName字段给出类所 在的文件名
	className  string								//className字段给出声明方法的类名
	methodName string								//methodName字段给出方 法名
	lineNumber int									//lineNumber字段给出帧正在执行哪行代码
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

//这个函数需要解释一下。由于栈顶两帧正在执行 fillInStackTrace(int)和fillInStackTrace()方法，所以需要跳过这两 帧。
//这两帧下面的几帧正在执行异常类的构造函数，所以也要跳 过，具体要跳过多少帧数则要看异常类的继承层次。
func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	//distanceToObject()函数计算所需跳过的帧数
	skip := distanceToObject(tObj.Class()) + 2
	//计算好需要跳过的帧之后，调用Thread结构体的GetFrames() 方法拿到完整的Java虚拟机栈，然后reslice一下就是真正需要的帧。
	//GetFrames()方法只是调用了Stack结构体的getFrames()方法
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		//createStackTraceElement()函数根据帧创建StackTraceElement实
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}
