package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"-"`    //忽略属性在json中打印
	Name     string `json:"name"` //{"name":"张三","Addr":"湖南"}
	sex      int    //首字母小写的，不能被json序列化
	tel      string
	Addr     string
	RealName string `json:"realName:string,omitempty"` //omitempty 如果是零值的时候，不做序列化
}

func main() {
	user := User{1, "张三", 1, "1898238", "湖南", "真实张三"}
	bytes, _ := json.Marshal(user)

	//小写的不可见，如果一个属性需要序列化，首字母需要大写
	fmt.Println(string(bytes)) //{"ID":1,"Name":"张三","Addr":"湖南"}

	var user02 User
	json.Unmarshal(bytes, &user02)
	fmt.Printf("%#v \n",user02)		//main.User{ID:0, Name:"张三", sex:0, tel:"", Addr:"湖南", RealName:"真实张三"}

}
