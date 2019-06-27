package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	list := os.Args // 获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage:src  dst")
		return
	}

	srcFile := list[1]
	dstFile := list[2]

	if srcFile == dstFile {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	//只读文件打开源文件
	sf,err1 := os.Open(srcFile)
	if err1 != nil {
		fmt.Println("err1 = ",err1)
		return
	}

	//新建目的文件
	ds,err2 := os.Create(dstFile)
	if err2 != nil {
		fmt.Println("err2 =" ,err2)
		return
	}

	// 操作完毕,需要关闭文件
	defer  sf.Close()
	defer  ds.Close()

	// 核心处理,从源文件读取内容,往目的文件写,读多少写多少
	buf := make([]byte,4*1024)
	for {
		n,err := sf.Read(buf)
		if err != nil{
			fmt.Println("err = ",err)
			if err == io.EOF{
				break
			}
		}
		// 往目的文件写,读多少,写多少
		ds.Write(buf[:n])
	}




}
