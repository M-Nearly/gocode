package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Readline(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Print("err = ", err)
		return
	}
	defer f.Close()

	//新建一个缓冲区,把内容先放在缓冲区
	r := bufio.NewReader(f)
	for {
		//遇到 \n 结束读取,但是 '\n'也读取进去了
		buf,err :=r.ReadBytes('\n')
		if err != nil{
			if err == io.EOF {//文件已经读完
				break
			}
			fmt.Println("err = ",err)
		}
		fmt.Printf("buf = #%s#\n",string(buf))
	}

}

func main() {
	path := "./model.txt"
	Readline(path)
}
