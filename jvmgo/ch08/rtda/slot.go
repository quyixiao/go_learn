package rtda

import "go_learn/jvmgo/ch08/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
