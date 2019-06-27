package main

import "fmt"
import "encoding/json"


// 成员变量名首字母必须大写
type IT struct {
	Company string `json:"company"` // 生成的json字符串会变成小写
	Subjects []string `json:"subjects"` // 二次编码
	IsOk bool	`json:"string"` //修改json的类型
	Price float64 `json:"-"` //此字段不会输出到屏幕
}
/*
type IT struct {
	Company string `json:"company"` // 生成的json字符串会变成小写
	Subjects []string `json:"subjects"` // 二次编码
	IsOk bool
	Price float64
}


*/

func main()  {
	//定义一个结构体变量,同事初始化
	s := IT{"itcast",[]string{"go","C++","Python","Test"},true,666.66}

	//编码,根据内容生成json文本
	//buf, err := json.Marshal(s)
	//if err != nil {
	//	fmt.Println("err =",err)
	//	return
	//}
	//fmt.Println("buf=",string(buf))
	//buf= {"Company":"itcast","Subjects":["go","C++","Python","Test"],"IsOk":true,"Price":666.66}


	//格式化输出json
	buf,err := json.MarshalIndent(s,""," ")
	if err != nil {
		fmt.Println("err=",err)
		return
	}
	fmt.Println("buf= ",string(buf))
	/*
	buf=  {
	 "Company": "itcast",
	 "Subjects": [
	  "go",
	  "C++",
	  "Python",
	  "Test"
	 ],
	 "IsOk": true,
	 "Price": 666.66
	}
	*/



}