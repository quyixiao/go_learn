package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T \n", now) //time.Time
	fmt.Printf("%v \n", now) //2021-11-07 16:45:30.621285 +0800 CST m=+0.000053567
	// 2000/01/02 08:10:10
	// 2006年
	// 01月
	// 02天
	// 24进制小时 15
	// 分钟 04
	// 秒 05

	fmt.Println(now.Format("2006/01/02 15:04:05")) //2021/11/07 16:45:30
	fmt.Println(now.Format("2006-01-02 15:04:05")) //2021-11-07 16:45:30
	fmt.Println(now.Format("2006-01-02"))          //2021-11-07
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Unix())     //1636274730
	fmt.Println(now.UnixNano()) //1636274730621285000

	xx, err := time.Parse("2006/01/02 03:04:05", "2006/01/02 03:04:05")

	fmt.Println(err, xx) //<nil> 2006-01-02 03:04:05 +0000 UTC

	t := time.Unix(0, 0)
	fmt.Println(t) //1970-01-01 08:00:00 +0800 CST

	d := now.Sub(t)
	fmt.Println(d)         //454520h51m59.778108s
	fmt.Printf("%T \n", d) //time.Duration

	// time.Second
	// time.Hour
	// time.Minute
	// timeMillisecond
	fmt.Println(time.Now()) //2021-11-07 16:54:33.066575 +0800 CST m=+0.000369435
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now()) //2021-11-07 16:54:38.070339 +0800 CST m=+5.004314271
	now3 := now.Add(time.Hour * 3)
	fmt.Println(now, now3) //2021-11-07 16:55:47.174147 +0800 CST m=+0.000093460 2021-11-07 19:55:47.174147 +0800 CST m=+10800.000093460

	d, err = time.ParseDuration("3h2m4s")
	fmt.Println(err, d) //<nil> 3h2m4s

	fmt.Println(d.Hours())   //3.0344444444444445
	fmt.Println(d.Minutes()) //182.06666666666666
	fmt.Println(d.Seconds()) //10924

}
