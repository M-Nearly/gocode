package main

import (
	"encoding/json"
	"fmt"
)

//注意大小写

func main() {
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
	//创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("tmp = ",tmp)
	fmt.Printf("tmp = %+v\n", m) // 打印结构

	// 取map中的字段
	//var str string
	//str = m["company"]
	//fmt.Println(str)  //err,无法转化 不能按照string的格式去取


	var str1 string
	//类型断言,值,他是value类型
	for key, value := range m {
		//fmt.Printf("%v ===> %v\n",key,value)
		switch data := value.(type) {
		case string:
			str1 = data
			fmt.Println("str = ",str1)
			fmt.Printf("map[%s]的值类型为string,value = %s\n", key, data)
		case bool:
			fmt.Printf("map[%s]的值类型为bool,value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64,value = %f\n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string,value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface{},value = %v\n", key, data)
		}

	}

}
