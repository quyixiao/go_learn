package main

import "fmt"

func main() {
	users := map[string]int{"zhangsan": 1, "lisi": 9, "wangwu": 10}
	keySlice := make([]string, len(users))
	valueSlice := []int{}
	i := 0
	for k, v := range users {
		keySlice[i] = k
		i++
		valueSlice = append(valueSlice, v)
	}
	for k,_ := range users{
		fmt.Println(k)
	}
	for _,v := range users{
		fmt.Println(v)
	}
	for v :=range users{
		fmt.Println(v)
	}
	for v:= range valueSlice{
		fmt.Println(v)
	}





}
