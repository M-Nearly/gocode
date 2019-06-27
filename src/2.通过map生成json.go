package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	//创建一个map
	m := make(map[string]interface{},4)
	m["company"] = "itcast"
	m["subjects"] = []string{"go","python","c++"}
	m["isok"] = true
	m["price"] = 123.123

	//编码json
	//res,err := json.Marshal(m)
	//格式化json
	res,err := json.MarshalIndent(m,"","	")
	if err != nil{
		fmt.Println("err = ",err)
		return
	}
	fmt.Println("res = ",string(res))
}