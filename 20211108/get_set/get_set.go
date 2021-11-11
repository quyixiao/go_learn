package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (user User) GetId() int {
	return user.id
}

func (user *User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

type Employee struct {
	User
	Salary float64
	name   string
}

func (employee Employee) GetName() string {
	return employee.name
}


func (employee * Employee) SetName(name string) {
	employee.name = name
}

func main() {
	u := User{
		id:   1,
		name: "张三",
	}
	u.SetId(10)
	fmt.Println(u.id)
	u.SetName("哈哈")
	fmt.Println(u.GetName()) //哈哈
	e := Employee{
		User:   User{1, "kk"},
		Salary: 1000,
		name:   "xiaoke",
	}
	fmt.Println(e.GetName())      		//xiaoke
	fmt.Println(e.User.GetName()) 		//kk
	e.SetName(" set userName ") 	//设置的是user的Name
	fmt.Println(e.GetName())      		// 获取到Employee的Name
	fmt.Println(e.User.GetName()) 		// 获取用户名
	e.SetName("哈哈")
	fmt.Println("----------",e.GetName())			// 哈哈


}
