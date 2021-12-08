package base

import "fmt"
import "strings"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"
//函数的前三行代码创建新的帧并推入Java虚拟机栈，剩下的代 码传递参数。
//首先，要确定方法的参数在局部 变量表中占用多少位置。注意，这个数量并不一定等于从Java代码 中看到的参数个数，原因有两个。
//第一，long和double类型的参数要占用两个位置。第二，对于实例方法，Java编译器会在参数列表的 前面添加一个参数，这个隐藏的参数就是this引用。
//假设实际的参 数占据n个位置，依次把这n个变量从调用者的操作数栈中弹出，放 进被调用方法的局部变量表中，参数传递就完成了。
//注意，在代码中，并没有对long和double类型做特别处理。因为 操作的是Slot结构体，所以这是没问题的。LocalVars结构体的 SetSlot()方法是新增的
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	//_logInvoke(callerFrame.Thread().StackDepth(), method)
	thread := invokerFrame.Thread()
	//如果要执行的是Java方法(而非本地方法)，下一步是给这个方
	//法创建一个新的帧并把它推到Java虚拟机栈顶。
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}

func _logInvoke(stackSize uint, method *heap.Method) {
	space := strings.Repeat(" ", int(stackSize))
	className := method.Class().Name()
	methodName := method.Name()
	fmt.Printf("[method]%v %v.%v()\n", space, className, methodName)
}
