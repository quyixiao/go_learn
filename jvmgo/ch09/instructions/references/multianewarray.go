package references

import "go_learn/jvmgo/ch09/instructions/base"
import "go_learn/jvmgo/ch09/rtda"
import "go_learn/jvmgo/ch09/rtda/heap"

// Create new multidimensional array
type MULTI_ANEW_ARRAY struct {
	index      uint16						//multianewarray指令的第一个操作数是个uint16索引，通过这个 索引可以从运行时常量池中找到一个类符号引用，解析这个引用就 可以得到多维数组类。
	dimensions uint8					//第二个操作数是个uint8整数，表示数组维 度。
}

//这两个操作数在字节码中紧跟在指令操作码后面，由 FetchOperands()方法读取
func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

//multianewarray指令还需要从操作数栈中弹出n个整数，分别代
//表每一个维度的数组长度。Execute()方法根据数组类、数组维度和 各个维度的数组长度创建多维数组
//这里提醒读者注意，在anewarray指令中，解析类符号引用后得 到的是数组元素的类，而这里解析出来的直接就是数组类。
func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(self.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		//popAndCheckCounts()函数从操作数栈中弹出n个int值，并且确保 它们都大于等于0。如果其中任何一个小于0，则抛出 NegativeArraySizeException异常。
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}

//newMultiArray()函数创建多维数组
func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
