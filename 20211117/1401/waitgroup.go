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

	n := 2
	group.Add(n)
	for i := 0; i < n; i++ {
		go func(id int) {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d:%d:%c\n", i, id, ch)
				runtime.Gosched()
			}
			// 让出CPU
			// time.Sleep(1 * time.Microsecond)
			group.Done()
		}(i)
	}

	group.Wait()
}
