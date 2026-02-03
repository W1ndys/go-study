package main

import "fmt"

func main() {
	languages := []string{"go", "py", "c"}
	fmt.Println("列表初始：", languages)
	fmt.Println("第一个：", languages[0])
	languages = append(languages, "rust")
	fmt.Println("追加后的列表：", languages)
	languages[1] = "c++"
	fmt.Println("修改后的列表：", languages)
}
