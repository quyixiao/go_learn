package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串=>其他类型
	// => bool
	if v, err := strconv.ParseBool("true"); err == nil {
		fmt.Println(v, err) // true <nil>
	}
	//转化成int
	if v, err := strconv.Atoi("1023"); err == nil {
		fmt.Println(v) //1023
	}

	if v,err := strconv.ParseInt("64",16,64); err == nil {
		fmt.Println(v)					// 100
	}

	//转化为64类型
	if v,err := strconv.ParseFloat("64.4 ",64); err == nil {
		fmt.Println(v)					// 100
	}




	//其他类型向string类型转化

	xx :=fmt.Sprintf("%d",12)
	fmt.Println(xx)


	yy :=fmt.Sprintf("%.3f",12.32389328)
	fmt.Println(yy) // 12.324

	fmt.Println(strconv.FormatBool(false)) // false
	fmt.Printf("%q \n",strconv.Itoa(12)) //"12"
	fmt.Printf("%q \n",strconv.FormatInt(12,16)) //"c"
	fmt.Printf("%q \n",strconv.FormatFloat(10.1 ,'E',-1,64)) //"1.01E+01" ,用科学计数法来表示

}
