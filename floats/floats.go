package main

import "fmt"

func main() {
	// float32 ,float64
	// 字面量： 有两种，一种是十进制表示法，一种是 科学计数法
	// 1 E N => 1
	// 1.9 E -1 => 0.19
	var height float64 = 1.68
	fmt.Printf("%T %f\n", height, height) // float64 1.680000



	var weight float64 = 0.19E1
	fmt.Println(weight) // 1.9




	// 操作
	// 算数运算( + ,- ,* ,/ , ++ )
	fmt.Println(1.2 + 1.3 )
	fmt.Println(1.1 -1.2 )
	fmt.Println( 1.1 / 1.2 )
	// fmt.Println(1.1. % 1.2 )
	weight ++
	fmt.Println( weight) // 2.9
	weight --
	fmt.Println(weight) // 1.9 浮点数在计算的时候有精度损失，因此一般不用来计算 （ == 计算）
	// 一般 只用来计算 （>= ,<= ,> ,< ）

	fmt.Println(1.2 - 1.2 <= 0.005 )				//true
	// 赋值（= ，+= ，-= ，/= ,*= ）
	height += 2.25
	fmt.Println(height)

	fmt.Printf("%T %T \n", 1.1 ,height) //float64 float64

	fmt.Printf("%5.2f",height)


	//复数
	ii := 1 + 2i
	fmt.Printf("%T %v\n",ii,ii)


}

