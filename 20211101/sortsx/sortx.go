package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 8,  10, 11}
	sort.Ints(nums)
	fmt.Println(nums)
	// 二分查找,在有序的数组中查找元素
	index := sort.SearchInts(nums, 5)
	fmt.Println(index) // 4
	// 先查找这个值，是不是等于num，如果不存在，则找到这个如果插入时的索引
	fmt.Println(sort.SearchInts(nums, 1000))

	if  nums[sort.SearchInts(nums, 9)] == 9 {		//如果不存在，直接打印不存在
		fmt.Println("存在 ")
	}else{
		fmt.Println("不存在 ")
	}
}
