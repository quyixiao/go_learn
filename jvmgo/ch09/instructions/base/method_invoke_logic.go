package base

import (
	"fmt"
	"go_learn/jvmgo/ch09/rtda"
	"go_learn/jvmgo/ch09/rtda/heap"
)


func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	//第一，long和double类型的参数要占用两个位置。
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			//第二，对于实例方法，Java编译器会在参数列表的 前面添加一个参数，这个隐藏的参数就是this引用。假设实际的参
			//数占据n个位置，依次把这n个变量从调用者的操作数栈中弹出，放 进被调用方法的局部变量表中，参数传递就完成了。
			slot := invokerFrame.OperandStack().PopSlot()
			//注意，在代码中，并没有对long和double类型做特别处理。因为 操作的是Slot结构体，所以这是没问题的。
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack!
	//由于Object类是其他所有类的超类，所以这会导致Java虚拟机 崩溃。解决办法是修改InvokeMethod()函数
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			//让它跳过所 有registerNatives()方法
			//如果遇到其他本地方法，直接调用panic()函数终止程序执行 即可。
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
