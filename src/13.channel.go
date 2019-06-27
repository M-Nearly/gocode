package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan string)
	defer fmt.Println("主协成也结束了")

	go func() {
		defer fmt.Println("zi xie cheng jieshu ")
		for i:=0; i <2;i++ {
			fmt.Println("子协成 i=",i)
			time.Sleep(time.Second)

		}
		 ch <- "我是子协成,要工作完毕"

	}()

	str := <- ch
	fmt.Println("str:",str)
}
