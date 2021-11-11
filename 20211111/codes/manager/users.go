package manager

import (
	"fmt"
	"sort"
	"strings"
)

type User struct {
	ID       int
	RealName string
}

//添加用户
func AddUser(users *[]User) {
start:
	var (
		id   int
		name string
	)
	fmt.Println("请输入姓名或退出(q/Q)：")
	fmt.Scan(&name)

	if name == "q" || name == "Q" {
		return
	}
	if exits(users, name) {
		fmt.Println(name, " 名字已经存在")
		goto start
	}
	id = GetmaxId(users)
	user := User{ID: id, RealName: name}
	*users = append(*users, user)
	for _, v := range *users {
		fmt.Printf("%#v \t ", v)
	}
	fmt.Println()
	goto start
}

func exits(users *[]User, name string) bool {
	for _, v := range *users {
		if strings.Contains(v.RealName, name) {
			return true
		}
	}
	return false
}

func GetmaxId(users *[]User) int {
	maxId := 1
	for _, v := range *users {
		if maxId < v.ID {
			maxId = v.ID
		}
	}
	maxId = maxId + 1
	return maxId
}

//更新用户
func UpdateUser(users *[]User) {
	fmt.Println("请输入要修改用户姓名:")
	var inputId string
	fmt.Scan(&inputId)
	xx := strings.Split(inputId, "_")
	fmt.Println(xx)
	for _, v := range *users {
		if strings.Contains(v.RealName, xx[0]) {
			v.RealName = xx[1]
			break
		}
	}
}

//删除用户
func DeleteUser(users *[]User) {

start:
	fmt.Println("请删除用户姓名:")
	var inputId string
	fmt.Scan(&inputId)
	if inputId == "q" || inputId == "Q" {
		return
	}
	if inputId == "all" {
		*users = []User{}
	} else {
		for k, v := range *users {
			if strings.Contains(v.RealName, inputId) {
				*users = append((*users)[:k], (*users)[k+1:]...)
				break
			}
		}
	}
	goto start
}

//
func QueryUser(users *[]User) {
start:
	fmt.Println("请输入要查询的用户姓名或退出(q/Q):")
	var inputId string
	fmt.Scan(&inputId)
	if inputId == "q" || inputId == "Q" {
		return
	}
	if inputId == "all" {
		sort.Slice(*users, func(i, j int) bool {
			return (*users)[i].RealName > (*users)[j].RealName
		})
		for _, v := range *users {
			fmt.Printf("%5d | %20s\n", v.ID, v.RealName)
		}
	} else {
		for _, v := range *users {
			if strings.Contains(v.RealName, inputId) {
				fmt.Printf("%5d | %20s \n", v.ID, v.RealName)
				break
			}
		}
	}
	goto start
}
