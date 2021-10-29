package main

import "fmt"

func main() {
	var name string = "k\tk" 	//"" 可解释的字符串
	var desc string = `我来自\t己中国` // `` 表示原生字符串
	// "" 和 `` 的区别， \r \n \t \b \f \v
	fmt.Println(name)
	fmt.Println(desc)


	// 操作
	// 算数运算符 :+(链接)
	fmt.Println(" 我叫" + "KK")

	// 关系运算(>,>= ,<= , == ,!= ,< )
	fmt.Println("a" == "b")  // false
	fmt.Println("a" != "b") //true
	fmt.Println("a"=="a") // true
	fmt.Println("aa" <= "bb")//true
	fmt.Println("aa">="ab") // false
	fmt.Println("ab">= "abb")// false
	fmt.Println("bb">="ba") // true


	s:= "我叫"
	s +="你是"
	fmt.Println(s) // 我叫你是


	// 字符串操作 ，字符串定义的内容，必须只能是ascii
	// 索引 0 -n - 1 (n 字符串的长度 )

	desc = "abcdef"
	fmt.Println(len(desc))

	fmt.Printf("%T %c \n", desc[2],desc[2]) // uint8 c
	// 切片  desc[start,end] 得到一个字符串
	fmt.Printf("%T %s \n",desc[0:2],desc[0:2])  //string ab
	desc = "我爱中国"

	fmt.Printf("%T %s \n",desc[0:2],desc[0:2])  // string �

	// 得到字符串的长度

	fmt.Println(len(desc)) // 12


}

