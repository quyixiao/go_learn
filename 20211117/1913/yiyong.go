package main

import "fmt"

func main() {
	n := 10
	channel := make(chan int, n)
	slice := make([]int,10)
	for i := 0 ;i < 10 ;i ++{
		select {

		case channel <- 0:
		case channel <- 1:
		case channel <- 3:
		case channel <- 4:
		case channel <- 5:
		}
		slice[i] = <-channel
	}
	fmt.Println(slice)
}
