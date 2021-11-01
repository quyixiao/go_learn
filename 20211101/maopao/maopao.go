package main

import "fmt"

func main() {
	heights := []int {10,6,7,9,5}
	//先把最高的人排到最后
	for i := 0 ;i < len(heights);i ++{
		for j := i + 1  ; j < len(heights)  ; j ++{
			if heights[i] > heights[j]{
				heights[i],heights[j] = heights[j],heights[i]
			}
		}
	}
	fmt.Println(heights)
}
