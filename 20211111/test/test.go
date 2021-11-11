package main

import (
	"fmt"
	"go_learn/20211111/codes/manager"
	"strconv"
)

func main() {
	listOfNumberStrings := []*string{}

	list := []*manager.User{}
	list = append(list,&manager.User{})

	for i := 0; i < 10; i++ {
		var numberString string
		numberString = fmt.Sprintf("#%s", strconv.Itoa(i))
		listOfNumberStrings = append(listOfNumberStrings, &numberString)
	}

	for _, n := range listOfNumberStrings {
		fmt.Printf("%s\n", *n)
	}


	fmt.Printf("%5d ,%20s \n",10,"idosi")
}
