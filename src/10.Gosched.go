package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {

			fmt.Println("go")
		}

	}()

	for i := 0; i < 2; i++ {
	// 让出时间片,先别别的协成执行,它执行完,在回到当前继续执行
		runtime.Gosched()
		fmt.Println("hello")
	}

}
