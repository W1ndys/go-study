package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 第一个返回值是字符在字符串的位置，第二个是字符的Unicode码点（ASCII码）
	for i, c := range "go" {
		fmt.Printf("i=%d, c=%d\n", i, c)
	}
}
