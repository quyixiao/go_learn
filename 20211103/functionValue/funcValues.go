package main

import "fmt"

func changeInt(a int) {
	a = 100
}

func changeSlice(s []int) {
	s[0] = 100
}

func changeIntAddress(n *int) {
	*n = 2

}

// 修改地址的地址对应的值
func changeIntAddressAddresss(n **int) {
	**n = 10
}

func main() {

	num := 1
	changeInt(num)
	fmt.Println(num) // 1  ,按值传递，是不会修改外面的值

	nums := []int{1, 2, 3}
	changeSlice(nums)
	fmt.Println(nums) // [100 2 3]       ，引用类型的数据是会修改外面的数据的

	changeIntAddress(&num)
	fmt.Println(num)					// 2

	num2 := &num
	changeIntAddressAddresss(&num2)
	fmt.Println(num)							// 10
}
