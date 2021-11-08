package usermanager

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"strconv"
	"strings"
)

//添加用户
func AddUser(users map[string]map[string]string) {
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
	id = GetmaxId(users)
	users[id] = map[string]string{
		"id":   id,
		"name": name,
	}
	fmt.Println(users)
	goto start

}

func GetmaxId(users map[string]map[string]string) string {
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

func InputString(prompt string) string {
	fmt.Println(prompt)
	var inputId string
	fmt.Scan(&inputId)
	return strings.Trim(inputId, " ")
}

//更新用户
func UpdateUser(users map[string]map[string]string) {
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
func DeleteUser(users map[string]map[string]string) {

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
func QueryUser(users map[string]map[string]string) {
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
	password = "202CB962AC59075B964B07152D234B70"
)

func getMd5Byte(bytes []byte) string {
	x := md5.Sum(bytes)
	return fmt.Sprintf("%X", x)
}



// 从命令行中输入值
func Auth() bool {
	fmt.Println("请输入密码：")
	for i := 0; i < maxAuth; i++ {
		bytes, _ := gopass.GetPasswd()
		if getMd5Byte(bytes) ==password{
			return true
		} else {
			fmt.Println("你的密码错误，请重新输入：")
		}
	}
	return false
}
