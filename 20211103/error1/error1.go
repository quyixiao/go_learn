package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("div by zero")
	}
	return a / b, nil
}

//返回值，怎样定义错误类型
// 怎么创建错误类型对应的值信息
func main() {
	//num := 1 / 0
	//fmt.Println(num) // ./error1.go:6:11: division by zero
	fmt.Println(div(1, 0))

	e := fmt.Errorf("error %s ","div by zero")
	fmt.Printf("%T  ,%v \n",e ,e)					// *errors.errorString  ,error div by zero

}
