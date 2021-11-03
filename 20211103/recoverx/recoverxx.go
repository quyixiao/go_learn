package main

import (
	"fmt"
)

func testxx() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	panic("error___________________")
}
func main() {
	err := testxx()
	fmt.Println(err)
	/*defer func() {
		fmt.Println()
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("恢复处理")

		}
	}()
	fmt.Println("main start ")

	panic("error xxx")
	fmt.Println("over")
*/
}
