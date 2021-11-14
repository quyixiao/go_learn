package manager

import (
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID       int
	RealName string
}

const (
	passwordfile = ".passwordfile"
	//user_path    = "users.gob"
	user_path = "users.csv"
)

func loadUserGob() *[]User {
	users := []User{}
	if file, err := os.Open(user_path); err == nil {
		defer file.Close()
		decodeer := gob.NewDecoder(file)
		decodeer.Decode(&users)
		fmt.Println("存储用户 ", users)
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("文件不存在", err)
		}
	}
	return &users
}

func loadUser() *[]User {
	users := []User{}
	if file, err := os.Open(user_path); err == nil {
		defer file.Close()
		reader := csv.NewReader(file)
		for {
			line, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					fmt.Println("发生错误 ", err)
				}
				break
			}
			id, _ := strconv.Atoi(line[0])
			user := User{
				ID:       id,
				RealName: line[1],
			}
			users = append(users, user)
		}
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("文件不存在", err)
		}
	}
	return &users
}

func saveUser(users *[]User) {
	// 重命名文件
	if _, err := os.Stat(user_path); err == nil {
		os.Rename(user_path, user_path+"_"+fmt.Sprintf("%d", time.Now().Unix())+".user.csv")
	}
	files, err := filepath.Glob("*.user.csv")
	// 删除文件,保存最近3个文件
	if err == nil {
		sort.Strings(files)
		fmt.Println(files)
		removeFiles := files[:len(files)-3]
		for _, removeFile := range removeFiles {
			fmt.Println("移除的文件是：" + removeFile)
			os.Remove(removeFile)
		}
	}
	if file, err := os.Create(user_path); err == nil {
		writer := csv.NewWriter(file)
		defer file.Close()
		for _, user := range *users {
			writer.Write([]string{strconv.Itoa(user.ID), user.RealName})
		}
		writer.Flush()
	}

}

func saveUserGob(users *[]User) {
	file, err := os.Create(user_path)
	if err == nil {
		defer file.Close()
		encoder := gob.NewEncoder(file)
		encoder.Encode(*users)
		fmt.Println("saveUser==========", users)
	}
}

//添加用户
func AddUser() {
start:
	users := loadUser()
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
	saveUser(users)
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
func UpdateUser() {
	users := loadUser()
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
	saveUser(users)
}

//删除用户
func DeleteUser() {

start:
	users := loadUser()
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
	saveUser(users)
	goto start
}

//
func QueryUser() {
start:
	users := loadUser()
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
