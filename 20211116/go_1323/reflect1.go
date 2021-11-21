package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
)

type User struct {
	id     int     `json:"id"`
	name   string  `json:"name"`
	Tel    string  `json:"addr"`
	Height float64 `json:"height"`
	Desc   *string `json:"desc"`
	Weight *int    `json:"weight"`
}

func NewUser(id int, name string, tel string, height float64, desc string, weight int) *User {
	return &User{id, name, tel, height, &desc, &weight}
}

func (user User) string() string {

	return fmt.Sprintf("User(id : %id ,name %s ,tel :%s ,height :%e ,desc :%s ,weight : %d )",
		user.id, user.name, user.Tel, user.Height, *user.Desc, *user.Weight)
}

func (user User) GetId() int {
	return user.id
}

func (user User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

// 定义接口Closer
type Closer interface {
	Close() error
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func (address Address) GetIp() string {
	return address.ip
}

func (address Address) GetPort() int {
	return address.port
}

type Connection struct {
	Address
	status int
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

func (conn *Connection) Send(msg string) error {
	fmt.Sprintf("发送消息给[%s:%d]:%s", conn.ip, conn.port, msg)
	return nil
}

func (conn *Connection) Close() error {
	fmt.Sprintf("关闭链接[%s:%d]", conn.ip, conn.port)
	return nil
}

// 打印reflect.Type类型变量的t信息
func displayType(t reflect.Type, tab string) {

	// 处理类型为nil
	if t == nil {
		fmt.Printf("<nil>")
		return
	}

	// 获取类型对应的枚举值 使用选择语句分别处理每种类型
	switch t.Kind() {
	case reflect.Int, reflect.Float32, reflect.Bool, reflect.String:
		//针对基本数据类型显示类名
		fmt.Printf("%s %s ", tab, t.Name())
	case reflect.Array, reflect.Slice:
		//针对数组和切片，直接打印Type对象
		fmt.Printf("%s %s ", tab, t)
	case reflect.Map:
		//对于映射类型打印键值对的Type对象
		fmt.Printf("%smap{\n", tab)
		fmt.Printf("%s\tKey:", tab)
		fmt.Printf("%s%s", tab, t.Key()) //获取键值的Type对象
		fmt.Println("")
		fmt.Printf("%s\tValue:", tab)
		fmt.Printf("%s%s", tab, t.Elem())
		fmt.Println()
		fmt.Printf("%s}", tab)
	case reflect.Func:
		// 针对函数打印参数和返回值，对于可变参数在最后一个参数之后添加 ....
		fmt.Printf("%s func ( ", tab)
		// 打印参数信息
		//获取参数数量并遍历
		for i := 0; i < t.NumIn(); i++ {
			fmt.Printf("%s", t.In(i)) //根据索引获取第i个参数的Type对象
			if i != t.NumIn()-1 {
				fmt.Printf(",")
			}

		}
		if t.IsVariadic() {
			fmt.Printf("....")
		}
		fmt.Printf(")")
		// 打印返回值信息
		if t.NumOut() > 0 {
			fmt.Printf("(")
			// 获取返回值数量并遍历
			for i := 0; i < t.NumOut(); i++ {
				fmt.Printf("%s", t.Out(i))
				if i != t.NumOut()-1 {
					fmt.Printf(",")
				}
			}
			fmt.Printf(")")
		}
		fmt.Printf("{}")

	case reflect.Struct:
		// 针对结构体显示结构体属性和方法
		fmt.Printf("%s type %s struct {\n", tab, t.Name())
		//获取属性数量并遍历
		fmt.Printf("%s \t Feilds(%d):\n", tab, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("%s \t %s \t %s \t %s ", tab, field.Name, field.Type, field.Tag) // 显示属性名，属性的Type对象和标签
			fmt.Printf(",\n")
		}

		// 获取方法的数量并遍历
		fmt.Printf("\n%s \tMethods(%d)\n", tab, t.NumMethod())
		for i := 0; i < t.NumMethod(); i++ {
			// 打印方法的信息
			displayMethod(t.Method(i), tab+"\t\t")
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Ptr:
		// 对于指针类型，递归分析其引用值
		fmt.Sprintf("%s *{\n", tab)
		displayType(t.Elem(), tab+"\t") //获取指针引用值，并递归调用displayType函数进行分析
		//打印指针变量的方法
		if t.NumMethod() > 0 { //获取方法的数量
			fmt.Printf("\n%s \t Methods (%d):\n", tab, t.NumMethod())
			for i := 0; i < t.NumMethod(); i++ {
				//打印方法信息
				displayMethod(t.Method(i), tab+"\t\t")
				if i != t.NumMethod()-1 {
					fmt.Printf(",\n")
				}
			}
		}
		fmt.Printf("\n%s}", tab)
	default:
		fmt.Printf("%s Unkonw[%s]", tab, t)
	}

}

//打印reflect.Method类型变量的method信息
func displayMethod(method reflect.Method, tab string) {
	// 获取方法的接收者类型
	t := method.Type
	fmt.Printf("%sfunc %s (", tab, method.Name) //显示方法名
	//打印参数信息
	//获取参数数量并遍历
	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf("%s", t.In(i)) //根据索引获取第i个参数的Type对象
		if i != t.NumIn()-1 {
			fmt.Printf(",")
		}
	}

	//打印可变参数信息
	fmt.Printf("...")
	//打印返回值信息
	fmt.Printf(")")

	if t.NumOut() > 0 {
		fmt.Printf("(")
		// 获取返回值数量并遍历
		for i := 0; i < t.NumOut(); i++ {
			fmt.Sprintf("%s", t.Out(i)) //根据索引获取第i个返回值的Type对象
			if i != t.NumOut()-1 {
				fmt.Printf(",")
			}
		}
		fmt.Printf(")")
	}
	fmt.Printf("{}")
}

// 打印reflect.Value类型的变量v的信息
func displayValue(value reflect.Value, tab string) {
	//获取对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Int:
		//针对整数类型，获取值对应的Type对象并使用Int函数转化为对应的基本类型，并使用strconv将转化为string类型。
		fmt.Printf("%s[%s]%s ", tab, value.Type(), strconv.FormatInt(value.Int(), 10))
	case reflect.Float64:
		//针对浮点类型，获取值对应的Type对象并使用Float函数转化为对应的基本类型，并使用strconv将转化为string类型
		fmt.Printf("%s [%s]%s ", tab, value.Type(), strconv.FormatFloat(value.Float(), 'E', -1, 64))
	case reflect.Bool:
		//针对布尔类型，获取值对应的Type对象并使用Bool函数转化为对应的基本类型，并使用strconv将转化为string类型
		fmt.Printf("%s[%s]%s", tab, value.Type(), strconv.FormatBool(value.Bool()))
	case reflect.String:
		//针对字符串类型，获取值对应的Type对象并打印存储
		fmt.Printf("%s[%s]%s", tab, value.Type(), value)
	case reflect.Array:
		//针对数组类型，获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", tab, value.Type())
		//获取数组的长度
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t") //根据索引获取数组的每个元素，并调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s} ", tab)
	case reflect.Slice:
		//针对切片类型，获取值对应的Type对象，长度和容量
		fmt.Printf("%s[%s](%d:%d) {\n", tab, value.Type(), value.Len(), value.Cap())
		//获取切片的长度
		for i := 0; i < value.Len(); i++ {
			displayValue(value.Index(i), tab+"\t") //根据索引获取数组中每个元素，并调用displayValue递归显示
			fmt.Printf(",\n")

		}
	case reflect.Map:
		//针对映射类型，获取值对应的Type对象
		fmt.Printf("%s[%s] {\n", tab, value.Type())
		//获取映射迭代对象遍历键值对
		iter := value.MapRange()
		for iter.Next() { //判断迭代对象是否到末尾
			displayValue(iter.Key(), tab+"\t") //根据从迭代对象获取当前键，并调用 displayValue递归显示
			fmt.Printf(":")
			displayValue(iter.Value(), "") //根据从迭代对象获取当前值，并调用 displayValue递归显示
			fmt.Printf(",\n")

		}
		fmt.Printf("%s}", tab)

	case reflect.Struct:
		//针对结构体类型，获取值对应的Type对象
		structType := value.Type()
		fmt.Printf("%s[%s] {\n", tab, structType)
		// 获取属性数量并遍历
		for i := 0; i < value.NumField(); i++ {
			structField := structType.Field(i)           //根据索引获取属性的类型
			field := value.Field(i)                      //根据索引获取属性的值
			fmt.Printf("%s\t%s:", tab, structField.Name) //打印类型名
			displayValue(field, tab+"\t")                //调用displayValue递归显示
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)
	case reflect.Ptr:
		//针对指针类型，获取值对应的Type对象
		fmt.Printf("%s[%s](\n", tab, value.Type())
		//获取指针的引用值，并递归调用displayValue函数递归分析显示
		displayValue(value.Elem(), tab+"\t")
		fmt.Printf("\n%s)", tab)

	default:
		fmt.Printf("%s Unkonw", tab)
	}

}

// 修改reflect.Value类型变量value的信息
func changeValue(value reflect.Value, path string) {
	//获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Int:
		//对于可设置的Int类型使用SetInt更新内存数据
		if value.CanSet() {
			fmt.Printf("Int CanSet :%s.%s \n", path, value.Type())
			value.SetInt(value.Int() + rand.Int63n(100))
		}
	case reflect.Float32:
		// 对于可设置
		if value.CanSet() {
			fmt.Printf("Float CanSet :%s.%s \n", path, value.Type())
			value.SetFloat(value.Float() + rand.Float64())
		}
	case reflect.Bool:
		//对于可设置的Float类型使用SetInt更新内存数据
		if value.CanSet() {
			fmt.Printf("Bool CanSet :%s ,%s \n", path, value.Type())
			value.SetBool(!value.Bool())
		}
	case reflect.String:
		//对于可设置的Float类型使用SetInt更新内存数据
		if value.CanSet() {
			fmt.Printf("Bool CanSet :%s.%s\n", path, value.Type())
			value.SetString("change:" + value.String())
		}
	case reflect.Array:
		//对于可设置的数组类型递归调用ChangeValue对每个元素的更新内存数据
		if value.CanSet() {
			fmt.Printf("Array CanSet :%s.%s \n", path, value.Type())
		}
		//获取数组的长度
		for i := 0; i < value.Len(); i++ {
			changeValue(value.Index(i), path+".array")
		}
	case reflect.Slice:
		//对于切片类型递归调用ChangeValue对每个元素的更新内存数据
		for i := 0; i < value.Len(); i++ {
			changeValue(value.Index(i), path+".slice") //根据索引获取切片的每个元素，并调用 ChangeValue递归参数
		}
	case reflect.Map:
		//对于映射类型调用SetMapIndex设置Key对应的值
		keys := value.MapKeys() //获取映射所有的key组成的Value切片
		for _, key := range keys {
			value.SetMapIndex(key, reflect.ValueOf("change:"+value.MapIndex(key).String()))
		}
	case reflect.Struct:
		if value.CanSet(){
			fmt.Printf("Struct CanSet:%s.%s\n",path,value.Type())
		}
		//对于结构休类型递归调用ChangeValue对每个元素更新内存数据
		for i:= 0 ;i < value.NumField();i ++{
			changeValue(value.Field(i),path +".struct") //根据索引获取结构体的每个属性，并调用changeValue递归修改
		}

	case reflect.Ptr:
		if value.CanSet(){
			fmt.Printf("CanSet :%s.%s \n",path,value.Type())
		}
		//对于指针类型，递归调用其其引用进行修改
		changeValue(value.Elem(),path + ".pointer")
	default:
		fmt.Printf("Unkonw[%#v]\n", value)
	}
}

// 定义string方法
func main() {
	vars := make([]interface{}, 0, 20)
	var intV int = 1
	var FloatV float32 = 3.14
	var boolV bool = true
	var stringV string = "吾日三省吾身，为人谋而不忠乎，与朋友交而不信乎？传不习乎？"
	var arrayV [5]int = [...]int{1, 2, 3, 4, 5}
	var sliceV []int = make([]int, 3, 5)
	var mapV map[string]string = map[string]string{"name": "kk"}
	var funcV1 func(...interface{}) error = func(x ...interface{}) error {
		fmt.Println(x...)
		return nil
	}
	var funcV2 func(string, int) *Connection = NewConnection
	var userV *User = NewUser(1, "kk", "15200000000", 1.68, "少年经不得顺境，中年经不得闲境，晚年经不得逆境", 72)
	var closerV Closer
	vars = append(vars, intV, &intV, FloatV, boolV, stringV, arrayV, sliceV, mapV, funcV1, funcV2, *userV, userV, closerV)

	for _, v := range vars {
		displayType(reflect.TypeOf(v), "")
		fmt.Println()
		fmt.Println("=======================")
	}

	fmt.Printf("********************************************************************")

	for _, v := range vars {
		displayValue(reflect.ValueOf(v), "")
		fmt.Println()
		fmt.Println("*****************************")
	}
	fmt.Printf("////////////////////////")
	// intV,floatV,boolV ,stringV,arrayV ,不可以修改
	// &intV,&floatV,&boolV,&stringV,&arrayV 可以修改，指针类型通过Elem获取的引用值可获取地址

	vars = append(vars,sliceV,mapV,*userV,userV,closerV)
	//sliceV,mapV 可以修改
	// userV 通过Elem获取的引用值可获取地址且公开的属性修改
	// * userV 公开的属性可修改（结构体通过Elem获取的引用值可获取地址）
	for _,v := range vars{
		fmt.Printf("-------------------------------------")
		fmt.Printf("%v\n",v)
		vv := reflect.ValueOf(v)
		changeValue(vv,"")
		fmt.Printf("%v\n",reflect.Indirect(vv))
	}
}
