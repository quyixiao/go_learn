package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s %#v  : 汪汪\n", dog.name, &dog)

}

func (dog *Dog) SetName(name string) {
	dog.name = name
	fmt.Printf("%#v \n", &dog)
}

type Cat struct {
	username string
	age      int
}

type Employ struct {
	Cat
}

func main() {
	m1 := Dog.Call
	// (*Dog).Call
	//方法表达式在赋值时，针对接收者为值类型的方法使用类型名或类型指针访问 （Go自动为指针变量生成隐匿的指针类型接收者方法 ），针对接收者
	// 为指针类型则使用类型指针访问，同时在调用 时需要传递对应的值 对象或指针对象
	m2 := (*Dog).Call
	// Dog.SetName
	fmt.Printf("%#v \n", m1) // (func(main.Dog))(0x108b740)
	fmt.Printf("%#v \n", m2) // (func(*main.Dog))(0x108b940)

	fmt.Printf("%T ,%T \n", m1, m2) //func(main.Dog) ,func(*main.Dog)
	dog := Dog{"豆豆"}

	m1(dog)
	dog.SetName("小黑1")
	m1(dog) //小黑1 &main.Dog{name:"小黑1"}  : 汪汪

	m2(&dog) //小黑1 &main.Dog{name:"小黑1"}  : 汪汪

	m3 := (*Dog).SetName

	m3(&dog, "小黄") //小黄 &main.Dog{name:"小黄"}  : 汪汪
	m1(dog)

	pdog := &Dog{"小白"}
	m1(*pdog) //小白 &main.Dog{name:"小白"}  : 汪汪

	b := Employ{
		Cat{
			username: "zhangsan",
		},
	}
	fmt.Printf(b.Cat.username)

}
