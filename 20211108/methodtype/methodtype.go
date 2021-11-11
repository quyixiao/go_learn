package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s %#v  : 汪汪\n", dog.name,&dog)

}

func (dog * Dog) CallPoint() {
	fmt.Printf("%s %#v  : 汪汪\n", dog.name,&dog)

}


func (dog *Dog) SetName(name string) {
	dog.name = name
	fmt.Printf("%#v \n",&dog)
}

func main() {
	dog := Dog{"豆豆"}
	m1 := dog.Call
	fmt.Printf("%T \n", m1)	//func()
	m1()							// 豆豆 : 汪汪

	dog.SetName("小黑")	//
	m1()						//豆豆 : 汪汪

	fmt.Println("============================")

	// 对于值接收者方法 者会拷备一份
	pdog := &Dog{"豆豆"}
	m2 := pdog.Call

	fmt.Printf("%T \n",m2)		//func()
	m2()						//豆豆 : 汪汪

	pdog.SetName("小黑")
	m2()					//豆豆 : 汪汪


	fmt.Println("============================")

	// 对于值接收者方法 者会拷备一份
	pdog1 := &Dog{"豆豆"}
	m3 := pdog1.CallPoint

	fmt.Printf("%T \n",m3)		//func()
	m3()						//豆豆 (**main.Dog)(0xc00000e040)  : 汪汪

	pdog1.SetName("小黑")
	m3()					//小黑 (**main.Dog)(0xc00000e050)  : 汪汪

}
