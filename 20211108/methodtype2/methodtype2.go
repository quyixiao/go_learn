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

func main() {

	m1 := Dog.Call
	m2 := (*Dog).Call
	fmt.Printf("%#v \n", m1)						// (func(main.Dog))(0x108b740)
	fmt.Printf("%#v \n", m2)						// (func(*main.Dog))(0x108b940)

	fmt.Printf("%T ,%T \n",m1,m2)					//func(main.Dog) ,func(*main.Dog)
	dog := Dog{"豆豆"}

	m1(dog)
	dog.SetName("小黑1")
	m1(dog)				//小黑1 &main.Dog{name:"小黑1"}  : 汪汪

	m2(&dog)			//小黑1 &main.Dog{name:"小黑1"}  : 汪汪

	m3 := (*Dog).SetName

	m3(&dog,"小黄")			//小黄 &main.Dog{name:"小黄"}  : 汪汪
	m1(dog)


	pdog := &Dog{"小白"}
	m1(*pdog)						//小白 &main.Dog{name:"小白"}  : 汪汪





}
