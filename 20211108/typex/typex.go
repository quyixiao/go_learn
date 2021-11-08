package main

import "fmt"

type Counter int

type User map[string]string

func main() {
	var counter Counter = 20
	counter += 10
	fmt.Println(counter) //30

	me := make(User)
	me["username"] = "zhangsan"
	me["age"] = "30"
	fmt.Println(me) //map[age:30 username:zhangsan]

	fmt.Printf("%T ,%T \n", counter, me)

	var callback = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, ":", v)
		}
	}

	callback("a", "a", "c")

	var counter2 int = 10
	fmt.Println(int(counter2) > counter2)
}
