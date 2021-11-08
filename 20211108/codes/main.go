package main

import (
	"fmt"
	"go_learn/20211108/codes/usermanager"
	"os"
)

func main() {
	if !usermanager.Auth() {
		fmt.Println("你的密码错误")
		return
	}
	callbacks := map[string]func(map[string]map[string]string){
		"1": usermanager.AddUser,
		"2": usermanager.UpdateUser,
		"3": usermanager.DeleteUser,
		"4": usermanager.QueryUser,
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
			os.Exit(500) //直接退出
		}
		callbacks[op](users)
	}

}
