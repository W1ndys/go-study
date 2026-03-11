package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("init slice", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set value", s)
	fmt.Println("get value", s[2])

	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("after append, s:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("after copy, c:", c)

	l := s[2:5]
	fmt.Println("s2-s5(not 5):", l)

	l = s[:5]
	fmt.Println("-s5(not 5)", l)

	l = s[2:]
	fmt.Println("s2-", l)

	t := []string{"g", "h", "i"}
	fmt.Println("声明并初始化：", t)

	// 多维数据结构，内部 slice 长度可以不一致
	towD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		towD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println("2d:", towD)
}
