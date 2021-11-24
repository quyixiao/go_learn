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

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
