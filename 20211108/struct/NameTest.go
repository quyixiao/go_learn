package main

import "fmt"

type User struct {
	Id   int
	Name string

}





type Employee struct {
	User
	Salary float64
	Name string
}


func main() {
	var e Employee
	e.Name = "张三"
	//如果只有一个属性，可以直接访问
	fmt.Printf("%#v \n",e)	//main.Employee{User:main.User{Id:0, Name:""}, Salary:0, Name:"张三"}

}

