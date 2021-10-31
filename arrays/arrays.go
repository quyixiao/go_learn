package main

import "fmt"

func main() {
	var nums [10]int
	fmt.Println(nums)         // [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("%q \n", nums) //['\x00' '\x00' '\x00' '\x00' '\x00' '\x00' '\x00' '\x00' '\x00' '\x00']
	fmt.Printf("%T \n", nums) // [10]int

	var t2 [5]bool
	var t3 [3]string
	fmt.Println(t2)         // [false false false false false]
	fmt.Println(t3)         //[  ] ,看到什么都没有打印，其实是5个空字符串
	fmt.Printf("%q \n", t3) // ["" "" ""] ,将对应类型的零值进行填充

	// 字面量,数组初始化
	nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums) // [1 2 3 4 5 6 7 8 9 10]
	// 给数组的第一个和最后一个赋值，[10 0 0 0 0 0 0 0 0 20]
	nums = [10]int{0: 10, 9: 20}
	fmt.Println(nums)
	// 可以根据... 去推算数组的长度
	nums = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //[1 2 3 4 5 6 7 8 9 10]
	fmt.Println(nums)
	// 定义一个数组长度为3的数组，可以不定义他的长度，可以用...来代替,可以根据数组的长度来推导出xx的长度
	xx := [...]int{1, 2, 3}
	fmt.Println(xx) //[1 2 3]

	//初始化数组的前三个元素的数组
	nums02 := [10]int{10, 20, 30}
	fmt.Println(nums02)                    // [10 20 30 0 0 0 0 0 0 0]
	fmt.Printf("%T %#v\n", nums02, nums02) //[10]int   [10]int{10, 20, 30, 0, 0, 0, 0, 0, 0, 0}

	//数组的声明也赋值
	nums03 := [...]int{1, 2}
	fmt.Printf("%T %#v\n", nums03, nums03) // [2]int       [2]int{1, 2}

	//获取数组的长度,数组的长度一旦定义了，就不能改变了
	fmt.Println("获取数组的长度",len(nums))					// 10
	// 遍历数组
	for i := 0; i < len(nums03); i++ {
		fmt.Println(nums03[i])
	}
	nums04 := [3]int{1, 2, 3}
	nums05 := [3]int{1, 2, 3}
	nums06 := [4]int{1, 2, 3, 4}
	nums07 := [4]int{1, 3, 4, 5}
	fmt.Println(nums04 == nums05) // true
	// invalid operation: nums04 == nums06 (mismatched types [3]int and [4]int) ,如果长度不相等，编译报错
	//fmt.Println(nums04 == nums06)
	fmt.Println(nums06 == nums07)						//false

	//索引,访问元素必需在索引范围之内
	fmt.Println("nums04的第一个元素 ：",nums04[0]) // nums04的第一个元素 ： 1

	//fmt.Println(nums04[-1]) Invalid array index '-1' (must be non-negative)
	nums04[0],nums04[1] = 1000,2000 //同时给数组中的元素赋值

	for index,value := range nums04{
		fmt.Println(index,value)
	}
	//0 1000
	//1 2000

	var value int 		//for里面的作用域和外部的作用域是不一样的
	//只想获取值，不想要index，那可以以一个空白标识符来标识，来丢弃不想要的值
	for _,value :=range nums04{
		fmt.Println(value)
	}
	fmt.Println(value)

	/*	var value int
	_,value = 1,4
	fmt.Println(value)*/
	// 切片
	fmt.Printf("%T \n",nums04[0:]) //[]int
	// [2000 3]
	fmt.Println(nums04[1:])
	//切片的容量，不能超过数组的长度
	//fmt.Println(nums04[1:4])
	fmt.Println("--------------------")
	// 最后一个参数是数组的容量
	fmt.Println(nums04[1:3:3])

	fmt.Printf("%v \n",nums04[1:3]) // [2000 3]

	// 多维数组
	//定义一个长度为2
	var marrays =  [3][2]int{{1,2},{3,4},{5,6}}
	fmt.Println(marrays)			// [[1 2] [3 4] [5 6]]
	fmt.Printf("%T\n",marrays)			// [3][2]int
	marrays[0] = [2]int{1,3}		// 修改第一维数组
	fmt.Println(marrays)
	marrays[1][1] = 1000
	fmt.Println(marrays)	//[[1 3] [3 1000] [5 6]]

	//定义三维数组,多维数组不能用可推导长度，多维数组的长度一定要是一一致的，数组的类型必需是一致的
	var marray3 = [4][3][2] int{{{1,2},{3,4},{5,6}},{{7,8},{9,10},{11,12}},{{13,14},{15,16},{17,18}},{{19,20},{21,22},{23,24}}}
	fmt.Println(marray3) // [[[1 2] [3 4] [5 6]] [[7 8] [9 10] [11 12]] [[13 14] [15 16] [17 18]] [[19 20] [21 22] [23 24]]]

}
