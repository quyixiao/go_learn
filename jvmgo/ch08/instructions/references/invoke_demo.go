package references

import "go_learn/jvmgo/ch08/instructions/base"
import "go_learn/jvmgo/ch08/rtda"
import "go_learn/jvmgo/ch08/rtda/heap"

// demo!
type INVOKE_XXX struct{ base.Index16Instruction }

func (self *INVOKE_XXX) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolved := resolveMethodRef(methodRef)
	checkResolvedMethod(resolved)
	toBeInvoked := findMethodToInvoke(methodRef)
	newFrame := frame.Thread().NewFrame(toBeInvoked)
	frame.Thread().PushFrame(newFrame)
	passArgs(frame, newFrame)
}

func resolveMethodRef(ref *heap.MethodRef) *heap.Method {
	// todo
	return nil
}
func checkResolvedMethod(method *heap.Method) {
	// todo
}
func findMethodToInvoke(ref *heap.MethodRef) *heap.Method {
	// todo
	return nil
}
func passArgs(caller, callee *rtda.Frame) {
	// todo
}
