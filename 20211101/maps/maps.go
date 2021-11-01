package main

import (
	"fmt"
)

func main() {
	// 映射的定义
	var scores map[string]int
	fmt.Printf("%T %#v\n", scores, scores) // map[string]int map[string]int(nil) 表示是一个nil的映射
	fmt.Println(scores == nil)             // true

	//字面量
	//scores = map[string]int{}
	scores = map[string]int{"武力": 10, "zhangsan": 4, "lisi": 20}

	fmt.Println(scores) // map[lisi:20 zhangsan:4 武力:10]

	xx := make(map[string]int)
	fmt.Printf("%T \n", xx) // map[string]int

	//map 增删改查
	// key
	// 从map中取值
	zhangsanScore := scores["zhangsan"]
	fmt.Println(zhangsanScore) // 4

	//在映射中，如果用key去访问，如果key不存在，则返回对应value的零值作为返回值
	fmt.Println(scores["xxx"]) // 用值的零值来作为返回值
	//第一个值是v ,第二个值是key是否存在
	v, ok := scores["xxx"]
	fmt.Println(ok, v) // false 0
	if ok {
		fmt.Println("xxx的成绩是", v)
	}

	//if包含初始化子语句
	if v, ok := scores["xxx"]; ok {
		fmt.Println(v)
	}
	scores["wangwu"] = 8
	fmt.Println(scores["wangwu"])
	scores["zhaoliu"] = 6
	fmt.Println(scores) // map[lisi:20 wangwu:8 zhangsan:4 zhaoliu:6 武力:10]
	//删除wangwu这个key
	delete(scores, "wangwu")
	fmt.Println(scores) //map[lisi:20 zhangsan:4 zhaoliu:6 武力:10]

	// 获取映射的长度
	fmt.Println(len(scores)) // 4

	// 遍历map,map是无序的
	for k, v := range scores {
		fmt.Println(k, v)
	}
	//武力 10
	//zhangsan 4
	//lisi 20
	//zhaoliu 6

	//不要key
	for _, v := range scores {
		fmt.Println(v)
	}
	//6
	//10
	//4
	//20

	// key 至少可以有== ,!=运算 ，bool ,整数，float , 字符串，数组，
	// 切片,映射是不能做key, 但可以做为value
	a := map[string][]int{"zhangsan": {1, 2, 3}, "wangwu": {1, 3, 5}}
	fmt.Println(a) // map[wangwu:[1 3 5] zhangsan:[1 2 3]]

	b := map[string]map[string]string{"zhangsan": {"key": "value"}, "lisi": {"hwhwh": "bbb"}}
	fmt.Println(b)               // map[lisi:map[hwhwh:bbb] zhangsan:map[key:value]]
	fmt.Printf("%T %#v\n", b, b) // map[string]map[string]string map[string]map[string]string{"lisi":map[string]string{"hwhwh":"bbb"}, "zhangsan":map[string]string{"key":"value"}}

	b["zhangsan"] = map[string]string{"aa": "bb","ccc":"ddd"}
	fmt.Println(b["zhangsan"])				// map[aa:bb ccc:ddd]


	// 删除映射中的值
	delete(b,"zhangsan") //
	fmt.Println(b)				// map[lisi:map[hwhwh:bbb]]


}
