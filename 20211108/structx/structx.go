package main

import "fmt"

type User struct {
	ID       int
	userName string
	age      int
	Addr     string
	remark   string
	tel      string
}

func main() {
	var user User
	fmt.Printf("%T \n", user) // main.User
	fmt.Printf("%#v\n", user) //main.User{ID:0, userName:"", age:0, Addr:"", remark:"", tel:""}

	var me2 = User{1, "张三", 10, "有道", "kwkw ", "182123455"}
	fmt.Printf("%#v\n", me2) //main.User{ID:1, userName:"张三", age:10, Addr:"有道", remark:"kwkw ", tel:"182123455"}

	var me3 = User{}
	fmt.Printf("%#v\n", me3)

	//指定属性名
	var me4 User = User{ID: 1, userName: "zhangsan", age: 10, Addr: "海南省", remark: "张", tel: "1845"}
	fmt.Printf("%#v\n", me4) //main.User{ID:1, userName:"zhangsan", age:10, Addr:"海南省", remark:"张", tel:"1845"}


	fmt.Printf("%T \n",&me3)	//*main.User 是一个指针类型


	var pointer *User
	fmt.Printf("%T \n",pointer)
	fmt.Printf("%#v \n",pointer)

	var point2 *User = &me2
	fmt.Printf("%#v \n",point2)
	var point3 *User = &User {}
	fmt.Printf("%#v \n",point3)

	//创建结构条的指针的时候，用new
	var pointer4 *User = new (User)
	fmt.Printf("%#v\n",pointer4)



}
