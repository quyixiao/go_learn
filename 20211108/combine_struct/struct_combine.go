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



type Company struct {
	ID int
	Name string
	Addr Address
	Salary float64
}



type Employee struct {
	User
	Company
	Salary float64
	Name string
}






func main() {

	var me01 User
	fmt.Printf("%#v \n", me01) //main.User{Id:0, Name:"", Addr:main.Address{Region:"", Streect:"", No:""}}

	user := User{
		Id:   1,
		Name: "zhangsan",
		Addr: Address{
			Region:  "北京",
			Streect: "天水街道",
			No:      "002",
		},
	}
	fmt.Printf("%#v \n", user) //main.User{Id:1, Name:"zhangsan", Addr:main.Address{Region:"北京", Streect:"和", No:"002"}}
	fmt.Println(user.Addr.Streect) //天水街道

	user.Addr.Streect = "沿河路"

	fmt.Println(user.Addr.Streect)					//沿河路


	var me Employee
	fmt.Printf("%T ,%#v \n",me)
	fmt.Printf("%T ,%#v \n",me.User.Name,me.User.Name)




}
