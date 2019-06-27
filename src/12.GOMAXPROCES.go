package main

import (
	"fmt"
	"runtime"
)

func main()  {
	n := runtime.GOMAXPROCS(2) // 已几核运算

	fmt.Println(n)
	for {
		go fmt.Println(1)

		fmt.Println(0)
	}
}
