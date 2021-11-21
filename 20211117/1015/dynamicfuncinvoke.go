package main

import (
	"fmt"
	"reflect"
)

func main() {
	//	定义func用户求和
	sum := func(paramters []reflect.Value) []reflect.Value {
		//若函数为不可变参数函数，则函数所有参数都包含在paraters中
		// 若函数为可变参数函数，则paramters最后一个元素为切片类型，包含调用函数传递的未对应的所有参数
		var total int64
		for _, paramter := range paramters[:len(paramters)-1] {
			total += paramter.Int()
		}
		//对最后一个参数进行特殊处理
		last := paramters[len(paramters)-1]
		switch last.Kind() {
		case reflect.Int:
			total += last.Int()
		case reflect.Slice:
			//若为切片，则为可变参数函数调用，遍历切片中的元素进行求和
			for i := 0; i < last.Len(); i++ {
				total += last.Index(i).Int()
			}
		}
		//返回包含总和初始化的Value切片，切片元素数量对应函数返回值的数量
		return []reflect.Value{reflect.ValueOf(total)}
	}

	var add2 func(int, int) int64
	ref := reflect.ValueOf(&add2).Elem()    //获取add2的引用
	fv := reflect.MakeFunc(ref.Type(), sum) //创建add2的函数值
	ref.Set(fv)
	//调用 add函数并打印结果
	fmt.Println(add2(1, 2))

	//定义匿名函数用于函数类型变量进行初始化
	//参数为函数类型的变量的指针
	makeFunc := func(fn interface{}) {
		ref := reflect.ValueOf(fn).Elem()       /// fn的引用
		fv := reflect.MakeFunc(ref.Type(), sum) // 根据fn的类型函数的值
		ref.Set(fv)
	}
	//定义函数变量
	var add3 func(int, int, int) int64
	var add4 func(int, int, int, int) int64

	//通过引用设置函数的变量值
	makeFunc(&add3)
	makeFunc(&add4)

	//调用函数
	fmt.Println(add3(1, 2, 3))    //6
	fmt.Println(add4(1, 2, 3, 4)) //10

	//定义可变参数类型的函数
	var add func(int, int, ...int) int64
	//通过引用设置函数的变量值
	makeFunc(&add)

	//调用函数
	fmt.Println(add(3, 4))    // 7
	fmt.Println(add(3, 4, 5)) // 12

}
