package main

import "fmt"

func main() {
	fmt.Println(add(1, 2)) //3
	fmt.Println(addN(1, 2, 3, 4, 5, 6)) // 9

	nums := []int{1,3,5,8}

	fmt.Println(nums) //[1 3 5 8]
	nums1 := append(nums[:1],5,8)// [1 5 8]
	fmt.Println("nums1",nums1) // nums1 [1 5 8]

	fmt.Println(nums) //[1 5 8 8]

	nums2 := append(nums[:1],nums[2:]...)
	fmt.Println(nums2) // [1 8 8]

	fmt.Println(nums) //[1 8 8 8]             可变参数底层公用
}
// 多个连续类型的参数
func add(a, b int) int {
	return a + b
}

// 可变参数只能定义一次，并且定义在最后
func calc(op string , a,b int ,args ... int ) int {
	switch op {
	case "add":
		return addN(a,b,args ...)
	}

	return 0
}
func addN(a, b int, args ...int) int {
	sum := a + b
	for _,v := range args {
		sum += v
	}
	return sum

}
