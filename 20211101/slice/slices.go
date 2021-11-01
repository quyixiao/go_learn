package main

import (
	"fmt"
)

func main() {
	//声明一个切片
	var nums []int
	fmt.Printf("%T %d %d \n", nums, len(nums), cap(nums)) //[]int 0 0

	var names []string
	fmt.Printf("%T %d %d %d \n", names, names == nil, len(names), cap(names)) // []string %!d(bool=true) 0 0

	// 字面量
	nums = []int{1, 2, 3}
	fmt.Println(nums) // [1 2 3]

	nums = []int{1, 2, 3, 4}                                //[1 2 3]
	fmt.Printf("%#v  %d %d \n", nums, len(nums), cap(nums)) // []int{1, 2, 3, 4} 4 4

	//通过数组切片赋值
	var arrays [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums = arrays[1:10]
	fmt.Println(nums) //通过数组切片，  [2 3 4 5 6 7 8 9 10]

	// make 函数,第一个参数是类型，第三个参数，是长度，第三个参数是容量
	nums = make([]int, 3)                                  //在底层初始化一个长度为3的数组，赋值给这个切片
	fmt.Printf("%#v %d %d \n", nums, len(nums), cap(nums)) // []int{0, 0, 0} 3 3
	fmt.Println(nums)

	nums = make([]int, 3, 5) //初始化前面三个参数是0，后面两个参数没有被初始化
	fmt.Println(nums)        // [0 0 0]

	// 元素的操作（增，删，改，查）
	fmt.Println(nums[0]) // 0
	fmt.Println(nums[1]) //0
	nums[2] = 4
	//	fmt.Println(nums[3])				// panic: runtime error: index out of range [3] with length 3
	//fmt.Println(nums[4])
	//nums[3]= 4						//
	// 为nums添加一个值
	nums = append(nums, 1)

	fmt.Printf("%#v %d %d \n", nums, len(nums), cap(nums)) // []int{0, 0, 4, 1} 4 5
	nums[3] = 10
	fmt.Println(nums) // [0 0 4 10]
	//	nums[5]=6							// panic: runtime error: index out of range [5] with length 4

	nums = append(nums, 1)
	nums = append(nums, 1)
	nums = append(nums, 1)
	fmt.Println(nums) // [0 0 4 10 1 1 1]

	// 切片遍历
	for _, value := range nums {
		fmt.Println(value)
	}

	// 切片操作
	nums = nums[1:5]
	fmt.Println(nums) //      [0 4 10 1]

	fmt.Println("%T \n", nums[1:5]) //  [4 10 1 1]

	nums = make([]int, 3, 10)
	n := nums[1:3:10] //设置的容量不能比原来的大
	//  长度是2  新的容量是 原来容量-start
	fmt.Printf("%T %#v  %d %d \n", n, n, len(n), cap(n)) // []int []int{0, 0}  2 9
	n = nums[2:3]
	// cap = src_cap - start
	fmt.Printf("%T %#v  %d %d \n", n, n, len(n), cap(n)) // []int []int{0}  1 8

	nums02 := nums[1:3]
	fmt.Println(nums02) // [0 0]
	nums02[0] = 1
	fmt.Println(nums02) // [1 0]
	// 删除
	// copy
	nums04 := []int{1, 2, 3}
	nums05 := []int{10, 20, 30, 40}
	//copy(nums05,nums04)
	//fmt.Println(nums05)					// [1 2 3 40]
	copy(nums04, nums05)
	fmt.Println(nums04) // [10 20 30] 复制不会扩容

	// 删除 第一个元素和最后一个元素
	nums06 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums06[1:])             // [2 3 4 5 6]
	fmt.Println(nums06[:len(nums06)-1]) //[1 2 3 4 5]
	//删除中间的一个元素
	copy(nums06[2:], nums06[3:])
	fmt.Println(nums06[:len(nums06)-1]) // [1 2 4 5 6]

	// 堆栈:每次添加在除尾，移除元素在队尾，（先进后出）

	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	fmt.Println(stack[:len(stack)-1])				// [1 2 3]

	//队列 ：每次添加都在队尾，移除元素在除头，先进先出

	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)
	fmt.Println(queue[0])             //1
	fmt.Println(queue[1:])            //[2 3 4]
	fmt.Println(queue[:len(queue)-1]) // [1 2 3]

	// 多维切片
	points := [][]int{}
	points02 := make([][]int,0)
	fmt.Printf("%T \n" ,points02)				// [][]int
	points = append(points,[]int{1,2,3})
	points = append(points,[]int{4,5,6})
	points = append(points,[]int{7,8,9,10,11,12})
	fmt.Println(points)
	fmt.Println(points[2][3])			// 10

	//数组的值的类型
	slice01 := []int{1,2,3}
	slice02 := slice01
	// 对切片改变，那另外的变量是会改变的
	slice02[0] = 100
	fmt.Println(slice01)			// [100 2 3]
	fmt.Println(slice02) //[100 2 3]

	//数组是一个值类型，在go语言中，所有的值类型，在赋值的时候，是值传递
	arrays01 := [3]int{1,2,3}
	arrays02 := arrays01
	arrays02[0] = 100
	fmt.Println(arrays01)			// [1 2 3]
	fmt.Println(arrays02)			//[100 2 3]


	// 切片的排序



}
