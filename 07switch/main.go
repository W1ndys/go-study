package main

import (
	"fmt"
	"time"
)

func main() {

	// 一个基本的switch
	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	// 逗号分割多个表达式，可选的default分支
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("今天是周末")
	default:
		fmt.Println("今天不是周末")
	}

	// 不带表达式的switch是实现if/else的一种形式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("现在是上午")
	default:
		fmt.Println("现在是下午")
	}

	// 类型开关
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("这是一个bool变量")
		case int:
			fmt.Println("这是一个int变量")
		default:
			fmt.Printf("未知类型：%T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(whatAmI)
}
