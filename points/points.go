package main

import "fmt"

func main() {
	A := 2
	B :=A
	A = 3
	fmt.Println(A,B) // 3 2

	C:=&A
	fmt.Printf("%T\n",C)			// *int

	var c *int = &A
	fmt.Println(c)		//如何修改


	fmt.Println(*c)

	*c = 4		//修改指向地址的值
	fmt.Println(*c,A)


	fmt.Printf("%T %T %p \n",c ,c,c ) // *int *int 0xc0000200a8





}
