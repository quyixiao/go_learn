package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	names := []string{"未子杰", "张三", "李四"}

	users := []map[string]string{{"name": "张三", "age": "30"}, {"name": "李四", "age": "40"}}

	bytes, err := json.Marshal(names)
	if err == nil {
		fmt.Println(string(bytes))
	}
	fmt.Println(users) //[map[age:30 name:张三] map[age:40 name:李四]]

	var names02 []string
	err = json.Unmarshal(bytes, &names02)

	fmt.Println(err)     //<nil>
	fmt.Println(names02) //[未子杰 张三 李四]

	//[
	//        {
	//                "age": "30",
	//                "name": "张三"
	//        },
	//        {
	//                "age": "40",
	//                "name": "李四"
	//        }
	//]
	userinfo, err := json.MarshalIndent(users,"","\t")
	if err == nil {
		fmt.Println(userinfo)         //[91 123 34 97 103 101 34 58 34 51 48 34 44 34 110 97 109 101 34 58 34 229 188 160 228 184 137 34 125 44 123 34 97 103 101 34 58 34 52 48 34 44 34 110 97 109 101 34 58 34 230 157 142 229 155 155 34 125 93]
		fmt.Println(string(userinfo)) //[{"age":"30","name":"张三"},{"age":"40","name":"李四"}]
	}
	var user02 []map[string]string
	err = json.Unmarshal(userinfo, &user02)
	fmt.Println(user02) //[map[age:30 name:张三] map[age:40 name:李四]]


	//判断一个字符串是不是json
	fmt.Println(json.Valid([]byte("[]x")))

}
