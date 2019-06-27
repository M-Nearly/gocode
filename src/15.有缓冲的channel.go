package main

import (
	"fmt"
	"time"
)

func main() {

	//创建一个有缓存的channel
	ch := make(chan int, 3)

	//len(ch) 缓冲区剩余数据个数, cap(ch) 缓冲区大小
	fmt.Println(len(ch), cap(ch))

	//新建协成
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			ch <- i
		}
	}()

	// 延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		//fmt.Println(i)
		num := <-ch
		fmt.Println("num=", num)
	}
}