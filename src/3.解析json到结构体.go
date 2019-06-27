package main

import (
	"encoding/json"
	"fmt"
)

//注意大小写
type IT struct {
	Company string `json:"company"` // 生成的json字符串会变成小写
	Subjects []string `json:"subjects"` // 二次编码
	IsOk bool	`json:"isok"` //修改json的类型
	Price float64 `json:"price"` //此字段不会输出到屏幕
}


func main()  {
	jsonBuf := `{

	"company": "itcast",
	"isok": true,
	"price": 123.123,
	"subjects": [
		"go",
		"python",
		"c++"
	]
}
`

	var tmp IT // 定义一个结构体
	err := json.Unmarshal([]byte(jsonBuf),&tmp)  //第二个参数要地址传递
	if err != nil{
		fmt.Println("err = ",err)
		return
	}
	//fmt.Println("tmp = ",tmp)
	fmt.Printf("tmp = %+v\n",tmp) // 打印结构

	//如果只想要其中的一个结构
	type IT2 struct {
		Subjects []string
	}
	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf),&tmp2)
	if err != nil{
		fmt.Println("err = ",err)
		return
	}
	//fmt.Println("tmp = ",tmp)
	fmt.Printf("tmp = %+v\n",tmp2) // 打印结构


}