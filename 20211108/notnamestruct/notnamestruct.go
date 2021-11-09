package main

import "fmt"

func main() {
	var me struct {
		id   int
		name string
	}

	fmt.Printf("%#v\n", me) // struct { id int; name string }{id:0, name:""}

	me2 := struct {
		Id   int
		name string
	}{1, "2"}

	fmt.Printf("%#v \n", me2)						// struct { Id int; name string }{Id:1, name:"2"}


}
