package main

import (
	"fmt"
	"time"
)

func newTask() {
	fmt.Println("这是新建的函数")
	time.Sleep(time.Second)

}
func main() {
	for {
		go newTask()

		fmt.Println("这是main函数")
		time.Sleep(time.Second)
	}

}
