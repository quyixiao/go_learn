package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("abc", "bad"))    // -1
	fmt.Println(strings.Contains("abc", "a"))     //true
	fmt.Println(strings.ContainsAny("abc", "cb")) //
	fmt.Println(strings.Count("abc", "a"))        //出现的次数
	fmt.Printf("%q \n", strings.Fields("abc cd")) //按空白字符分隔开 ，["abc" "cd"]

	fmt.Println(strings.HasPrefix("abc", "a")) // 以什么开头 true
	fmt.Println(strings.HasSuffix("bcd", "d")) // 以什么结尾

	fmt.Println(strings.Index("abc", "c"))      // 2 返回第一次出现的位置
	fmt.Println(strings.LastIndex("abac", "a")) // 2 返回最后一次出现的位置

	fmt.Println(strings.Split("abcdef", "e")) // [abcd f]
	a := []string{"a", "b", "c"}
	fmt.Println(strings.Join(a,"_")) // a_b_c				第一个参数是一个切片，第二个参数是连接的_


	fmt.Println(strings.Repeat("abc",3))					// abcabcabc
	fmt.Println(strings.ReplaceAll("abcabcabc","ab","x"))	// xcxcxc ,将ab 替换成x
	fmt.Println(strings.Replace("abcabcabc","ab","x",1))	// xcabcabc ,将ab 替换成x,只替换一个
	fmt.Println(strings.Replace("abcabcabc","ab","x",-1))	// xcxcxc ,将ab 替换成x,-1，表示替换所有
	fmt.Println(strings.ToLower("ABC"))					// abc
	fmt.Println(strings.ToUpper("abc")) //ABC 全部大写
	fmt.Println(strings.ToTitle("abc")) //ABC
	fmt.Println(strings.Trim("abc","c"))	//ab
	fmt.Println(strings.TrimSpace("abcxx iewoiew \t"))				//abcxx iewoiew 去前后的空白字符





}
