package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//添加用户
func addx(users map[string]map[string]string) {
start:
	var (
		id   string
		name string
	)
	fmt.Println("请输入姓名或退出(q/Q)：")
	fmt.Scan(&name)
	if name == "q" || name == "Q" {
		return
	}
	id = getmaxId(users)
	users[id] = map[string]string{
		"id":   id,
		"name": name,
	}
	fmt.Println(users)
	goto start

}

func getmaxId(users map[string]map[string]string) string {
	maxId := 1
	for k, _ := range users {
		if i, err := strconv.Atoi(k); err == nil {
			if maxId < i {
				maxId = i
			}
		}
	}
	maxId = maxId + 1
	return strconv.Itoa(maxId)
}

func inputString(prompt string) string {
	fmt.Println(prompt)
	var inputId string
	fmt.Scan(&inputId)
	return strings.Trim(inputId, " ")
}

//更新用户
func updatex(users map[string]map[string]string) {
	fmt.Println("请输入要修改用户姓名:")
	var inputId string
	fmt.Scan(&inputId)
	xx := strings.Split(inputId, "_")
	fmt.Println(xx)
	for k, v := range users {
		for _, name := range v {
			if strings.Contains(name, xx[0]) {
				users[k]["name"] = xx[1]
				break
			}
		}
	}
}

//删除用户
func deletex(users map[string]map[string]string) {
start:
	fmt.Println("请删除用户姓名:")
	var inputId string
	fmt.Scan(&inputId)
	if inputId == "q" || inputId == "Q" {
		return
	}
	if inputId == "all" {
		for k, _ := range users {
			delete(users, k)
		}
	} else {
		for k, v := range users {
			for _, name := range v {
				if strings.Contains(name, inputId) {
					delete(users, k)
					break
				}
			}
		}
	}
	goto start
}

//
func queryx(users map[string]map[string]string) {
start:
	fmt.Println("请输入要查询的用户姓名或退出(q/Q):")
	var inputId string
	fmt.Scan(&inputId)
	if inputId == "q" || inputId == "Q" {
		return
	}
	if inputId == "all" {
		for _, v := range users {
			fmt.Printf("%5s|%20s\n", v["id"], v["name"])
		}
	} else {
		for _, v := range users {
			for _, name := range v {
				if strings.Contains(name, inputId) {
					fmt.Printf("%5s|%20s\n", v["id"], v["name"])
					break
				}
			}
		}
	}
	goto start
}

const (
	maxAuth  = 3
	password = "132"
)

// 从命令行中输入值
func auth() bool {
	var input string
	fmt.Println("请输入密码：")
	for i := 0; i < maxAuth; i++ {
		fmt.Scan(&input)
		if input == password {
			return true
		} else {
			fmt.Println("你的密码错误，请重新输入：")
		}
	}
	return false
}

func main() {
	if !auth() {
		fmt.Println("你的密码错误")
		return
	}
	callbacks := map[string]func(map[string]map[string]string){
		"1": addx,
		"2": updatex,
		"3": deletex,
		"4": queryx,
	}
	fmt.Println(callbacks)
	//存储用户信息
	users := make(map[string]map[string]string)
	fmt.Println("欢迎使用用户管理系统：")
	//id := 0
	for {
		var op string
		fmt.Println(`请输入指令：
1.新建用户 2.修改用户 3.删除用户 4.查询用户 q.退出`)
		fmt.Scan(&op)
		if "q" == op {
			//return
			os.Exit(500)						//直接退出
		}

		callbacks[op](users)
		/*switch op {
		case "1":
			id++
			addx(&id, users)
			break
		case "2":
			updatex(users)
			break
		case "3":
			deletex(users)
			break
		case "4":
			queryx(users)
			break
		case "q":
			return
		default:
			break
		}*/
	}
}
