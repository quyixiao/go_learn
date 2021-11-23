package rtda

import "go_learn/jvmgo/ch09/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
