package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 1
	fmt.Printf("%T \n", i)				//int
	var tp reflect.Type = reflect.TypeOf(i)
	fmt.Println(tp.Name())						//int
	fmt.Println(tp.Kind())							//int类型
	fmt.Println(tp.PkgPath())				//int

	

}
