package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID   int
	Name string
}

func main() {

	list := [][2]int{{1, 3}, {5, 9}, {4, 3}, {6, 2}}

	//使用数组中第二个元素比较大小,元素比较大小进行排序
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] < list[j][1]
	})
	fmt.Println(list) //从小到大开始排序
	users := []User{
		{ID: 1, Name: "zhangsan"},
		{ID: 2, Name: "lisi"},
		{ID: 4, Name: "wangwu"},
		{ID: 3, Name: "zhaoliu"},
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID //从小到大开始排序
	})
	fmt.Println(users) //[{1 zhangsan} {2 lisi} {3 zhaoliu} {4 wangwu}]

}
