package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{4, 5, 7, 8, 6}
	sort.Ints(nums)
	//对nums排序
	fmt.Println(nums) // [4 5 6 7 8]

	names := []string{"test","name","zhangsan","lisi"}
	sort.Strings(names)//对字符串进行排序
	fmt.Println(names) //[lisi name test zhangsan]


	hight := []float64{1.2,1.0,2.5,3.2,2.3}
	sort.Float64s(hight) //对float类型进行排序
	fmt.Println(hight) // [1 1.2 2.3 2.5 3.2]

}
