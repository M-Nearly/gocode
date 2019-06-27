package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("ccccccccc")

	//return // 终止此函数
	runtime.Goexit() // 终止所在的协成  会退出 整个匿名函数的协成
	fmt.Println("dddddddddddddddddd")

}

func main() {
	go func() {
		fmt.Println("aaaaaaaaaa")

		test()

		fmt.Println("bbbbbbbbbb")
	}()

	for {}
}
