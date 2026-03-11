package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("初始化数组 a：", a)

	// 赋值
	a[4] = 100
	fmt.Println("赋值之后", a)
	fmt.Println("a[4]=", a[4])

	fmt.Println("len=", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("初始化数组 b:", b)

	// 二维数组
	var towb [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			towb[i][j] = i + j
		}
	}

	fmt.Println("towb=", towb)

}
