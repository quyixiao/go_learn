package base

import "go_learn/jvmgo/ch08/rtda"
import "go_learn/jvmgo/ch08/rtda/heap"

// jvms 5.5
func InitClass(thread *rtda.Thread, class *heap.Class) {
	//nitClass()函数先调用StartInit()方法把类的initStarted状态设 置成true以免进入死循环
	class.StartInit()
	//然后调用scheduleClinit()函数准备执行 类的初始化方法
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	//类初始化方法没有参数，所以不需要传递参数。
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	//注意，这里有意使用了scheduleClinit这个函数名而非 invokeClinit，因为有可能要先执行超类的初始化方法，
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			//如果超类的初始化还没有开始，就递归调用InitClass()函数执 行超类的初始化方法，这样可以保证超类的初始化方法对应的帧在
			//子类上面，使超类初始化方法先于子类执行。
			InitClass(thread, superClass)
		}
	}
}


