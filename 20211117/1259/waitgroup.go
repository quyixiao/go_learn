package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printChars(prefix string, group *sync.WaitGroup) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", prefix, ch)
		runtime.Gosched()

	}
	// 让出CPU
	// time.Sleep(1 * time.Microsecond)
	group.Done()
}

func main() {
	group := &sync.WaitGroup{}

	group.Add(3)
	go printChars("1", group)
	go printChars("2", group)

	printChars("3",group)

	group.Wait()
}
