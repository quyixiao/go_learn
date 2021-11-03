package main

import "fmt"

func main() {
	array:= [3]string{"A","B","C"}
	slice := []string{"A","B","C"}


	arrayA := array
	sliceA := slice

	arrayA[0] = "D"
	sliceA[0] = "D"
	fmt.Println(array ,arrayA == array)				//值类型
	//值类型，引用类型
	// 将变量赋值给新的一个变量，如果对变量有影响，引用类型，如果对变量没有影响，就是值类型
	fmt.Println(slice)


	// int ,bool,float32 ,float64 ,array ,slice ,map
	//  int ,bool,float32 ,float64 ,array，结构体 值类型
	// slice ,map ,接口，引用类型
	// 指针是一个值类型
	m := map[string]string{}
	mA := m
	mA["kk"] = "张天"
	fmt.Println(m) // map[kk:张天]

	//通过指针来改值
	age :=30
	b := &age
	*b = 20
	fmt.Println(age,*b)


	fmt.Printf("%p %p \n", & age ,& b )// 0xc0000b2020 0xc0000ac020
	fmt.Printf("%p %p \n", slice ,sliceA )	// 0xc000098180 0xc000098180
	fmt.Printf("%p %p \n", & slice[0] ,& sliceA[0] )// 0xc000098180 0xc000098180
	fmt.Printf("%p %p \n", & array,&arrayA ) 	// 0xc00006e180 0xc00006e1e0


}













