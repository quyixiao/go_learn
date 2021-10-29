package main

import "fmt"

func main() {
	// 整数类型
	// 标识符： int /int* /unit/unit*/unitptr/byte/rune
	// 字面量 ： 十进制，八进制 0777  = 7 * 8^0 + 7 * 8 ^1 + 7 * 8^2 ， 十六进制
	var age int = 31
	fmt.Println(age)
	fmt.Printf("%T %d ",age,age)
	fmt.Println(0777)
	// 操作
	// 算术运算（+，-，*, / ,% ）
	fmt.Println( 1 + 2 )
	fmt.Println( 3 - 10 )
	fmt.Println(3 *5 )
	fmt.Println(1 / 2 )
	fmt.Println(9 % 2 )

	age ++


	fmt.Println(age)

	age --
	fmt.Println(age)



	// 关系运算
	fmt.Println(2 == 3 )
	fmt.Println( 2 != 3 )
	fmt.Println(2 > 3 )
	fmt.Println(2 >= 3 )
	fmt.Println(2 < 3 )
	fmt.Println(2 <= 3 )


	// 位运算 ,计算机中的所有数字都按位存储的， 10 >= 2
	// & | ^ << >> &^

	// 十进制 => 2  7 /2  => 1  3 /2 => 1   1 /2 1
	// 2 => 0010
	// 7 & 2 => 0111 &  0010 => 0010
	// 7| 2 => 0111 & 0010 => 0111
	// 7 ^ 2 => 0111 ^ 0010 => 0101
	// 2 << 1 =>
	// 2 >> 1
	fmt.Println( 7 & 2 )
	fmt.Println( 7 | 2 )
	fmt.Println(7 ^ 2 )
	fmt.Println( 2 << 1 )		// 0010 << 1  => 0100
	fmt.Println(2 >> 1 )		// 0010 >> 1 => 0001
	fmt.Println(7 &^2 ) 	//



	// 赋值运算（ = ,+= ,-= ,*= ,/= ,%=, |= ,^= ,<<= ,>>= , &^=  ）
	// a+=b => a= a + b ,
	//
	age = 1
	age +=3
	fmt.Println(age)

	// int / uint /byte ,rune /int*
	var intA int = 10
	var unitB uint = 3
	// fmt.Println(intA + unitB ) 两种不同的数据类型，需要用户自己强制类型转换
	fmt.Println(intA + ((int)(unitB)))
	fmt.Println(((uint)(intA)) + unitB)

	// 从大往小转换，可能出现溢出的现象
	var intC int = 0XFFFF
	fmt.Println(intC,uint8(intC),int8(intC))


	// byte,rune
	var a byte = 'A'
	var w rune = '中'
	print(a , "-----------",w)
	fmt.Println("\n")

	fmt.Printf("%T %d %b %o %x \n",age ,age ,age ,age ,age ) // int 4 100 4 4
	fmt.Printf("%T %c \n", a ,a ) //uint8 A
	fmt.Printf("%T %q \n", w ,w) // int32 '中'
	fmt.Printf("%10d \n",age) //         4
	fmt.Printf("%010d \n",age)//0000000004










}
