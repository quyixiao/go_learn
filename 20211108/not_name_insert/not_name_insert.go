package main

import "fmt"

type Address struct {
	Region  string
	Streect string
	No      string
}

type User struct {
	Id   int
	Name string
	Addr Address
}

type Employee struct {
	User
	Salary float64
}

func main() {

	var me Employee
	fmt.Printf("%T,%#v", me, me)//main.Employee,main.Employee{User:main.User{Id:0, Name:"", Addr:main.Address{Region:"", Streect:"", No:""}}, Salary:0}



}
