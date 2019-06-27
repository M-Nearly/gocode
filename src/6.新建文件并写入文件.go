package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Println("n = ", n)
	}
}

func ReadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("f =", f)
	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) // 按照固定大小读取文件内容

	// n 代表从文件读取内容的长度
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {  // 文件出错,同时没有到结尾
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
}


func main() {
	path := "./model.txt"
	//WriteFile(path)
	ReadFile(path)
}
