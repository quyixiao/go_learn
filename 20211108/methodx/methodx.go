package main

import "fmt"

type Dog struct {
	Name string
}

func (dog Dog) Call(){
	fmt.Printf("%s :汪汪\n",dog.Name)
}


func (dog Dog ) SetName(name string ) {
	dog.Name = name
}



func (dog * Dog ) PsetName(name string ){
	dog.Name = name
}



func test(dog Dog ){

}

func main() {

	dog := Dog{"豆豆"}
	dog.Call()				//豆豆 :汪汪
	//dog.Name = "张三"
	//dog.Call()
	dog.SetName("哈哈")
	dog.Call()						//豆豆 :汪汪

	// (&dog).PsetName("小黑")取引用
	(*&dog).PsetName("小黑")
	dog.Call()						//小黑 :汪汪

	pdog := &Dog{"豆豆"}
	(*pdog).Call()


	// (*pdog).Call()解引用
	pdog.PsetName("小黑")
	(*pdog).Call()		//自动的解引用，语法糖


	test(*pdog)


//	var pdog2 *Dog
//	fmt.Println(pdog2.Call)					///Users/quyixiao/go/src/go_learn/20211108/methodx/methodx.go:56 +0x229


}
