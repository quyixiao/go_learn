package reserved

import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/native"
import _ "go_learn/jvmgo/ch11/native/java/io"
import _ "go_learn/jvmgo/ch11/native/java/lang"
import _ "go_learn/jvmgo/ch11/native/java/security"
import _ "go_learn/jvmgo/ch11/native/java/util/concurrent/atomic"
import _ "go_learn/jvmgo/ch11/native/sun/io"
import _ "go_learn/jvmgo/ch11/native/sun/misc"
import _ "go_learn/jvmgo/ch11/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	//根据类名、方法名和方法描述符从本地方法注册表中查找本 地方法实现，如果找不到，则抛出UnsatisfiedLinkError异常
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	//否则直 接调用本地方法
	nativeMethod(frame)
}
