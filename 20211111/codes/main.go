package main

import (
	"fmt"
	"go_learn/20211111/codes/manager"
	"os"
)

func main() {
	callbacks := map[string]func(users *[]manager.User){
		"1": manager.AddUser,
		"2": manager.UpdateUser,
		"3": manager.DeleteUser,
		"4": manager.QueryUser,
	}
	fmt.Println(callbacks)
	//存储用户信息
	users := []manager.User{}
	fmt.Println("欢迎使用用户管理系统：")
	//id := 0
	for {
		var op string
		fmt.Println(`请输入指令：
1.新建用户 2.修改用户 3.删除用户 4.查询用户 q.退出`)
		fmt.Scan(&op)
		if "q" == op {
			os.Exit(500) //直接退出
		}
		callbacks[op](&users)
	}

}
