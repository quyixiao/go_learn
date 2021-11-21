package main

import (
	"fmt"
	"reflect"
)

//处理reflect.Value类变量 v中所有的函数
func callValue(value reflect.Value) {

	fmt.Println("-----------------", value)
	// 获取值对应的枚举类型使用选择语句分别处理每种类型
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		//针对数组/切片类型，对数组/切片中元素进行递归处理
		for i := 0; i < value.Len(); i++ {
			callValue(value.Index(i)) //根据索引获取数组中的每个元素，并调用callValue递归调用

		}
	case reflect.Map:
		//针对映射类型，对映射值进行递归处理
		iter := value.MapRange()
		for iter.Next() {
			callValue(iter.Value()) //根据切片获取数组每个元素，并调用callValue递归调用
		}
	case reflect.Struct:
		//针对结构体类型，执行所有的方法
		for i := 0; i < value.NumMethod(); i++ {
			callMethod(value.Method(i), value.Elem().Type(), value.Type().Method(i))
		}
	case reflect.Func:
		//针对方法类型，执行方法
		callFunc(value)
	default:
		fmt.Printf("Unkown [%#v ]\n", value)

	}

}

//调用函数
func callFunc(f reflect.Value) {
	//获取函数的类型并打印
	ftype := f.Type()
	fmt.Printf("%s \n", ftype.String())
	//组装函数参数
	parameters := make([]reflect.Value, 0)
	for i := 0; i < ftype.NumIn(); i++ {
		parameters = append(parameters, reflect.Zero(ftype.In(i))) //通过reflect.Zeror创建参数同类型零值
	}
	//调用函数并打印执行结果
	if ftype.IsVariadic() {
		fmt.Printf("%#v", f.CallSlice(parameters))
	} else {
		fmt.Println(f.Call(parameters))
	}

}

//调用方法
func callMethod(f reflect.Value, t reflect.Type, m reflect.Method) {
	//获取函数并打印
	ftype := f.Type()
	fmt.Printf("method :%s.%s => %s \n", t.Name(), m.Name, ftype.String())
	//组装函数参数
	parameters := make([]reflect.Value, 0)
	for i := 0; i < ftype.NumIn(); i++ {
		parameters = append(parameters, reflect.New(ftype.In(i)).Elem()) //通过reflect.New创建参数同类型
	}

	//调用函数并打印执行结果
	if ftype.IsVariadic() {
		fmt.Println(f.CallSlice(parameters))
	} else {
		fmt.Println(f.Call(parameters))
	}
}

type Connection struct {
	status  int
	Address Address
}

type Address struct {
	ip   string "json:ip"
	port int    "json:port"
}

func NewConnection(ip string, port int) *Connection {
	return &Connection{Address: Address{ip, port}}
}

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

func main() {
	vars := make([]interface{}, 0, 20)
	var funcs []func() = make([]func(), 0)
	funcs = append(funcs,
		func() {
			fmt.Println(1)
		},
		func() {
			fmt.Println(2)
		},
		func() {
			fmt.Println(3)
		},
	)
	var funcV1 func(...interface{}) error = func(x ...interface{}) error {
		fmt.Println(x...)
		return nil
	}
	var funcV2 func(string, int) *Connection = NewConnection
	var userV *User = NewUser(1, "kk", "1520000", 1.68, "少年不知愁滋味", 72)
	vars = append(vars, funcs, funcV1, funcV2, *userV, userV)
	for _, v := range vars {
		callValue(reflect.ValueOf(v))
	}
}
