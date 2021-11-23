package heap

func (self *Class) IsArray() bool {
	return self.name[0] == '['
}

//Class结构体的ComponentClass()方法返回数组类的元素
//ComponentClass()方法先根据数组类名推测出数组元素类名， 然后用类加载器加载元素类即可。
func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}

//NewArray()方法专门用来创建数组对象。如果类并不是数组 类，就调用panic()函数终止程序执行，否则根据数组类型创建数组 对象。注意:布尔数组是使用字节数组来表示的。
func (self *Class) NewArray(count uint) *Object {
	//其中定义IsArray()方法
	if !self.IsArray() {
		panic("Not array class: " + self.name)
	}
	switch self.Name() {
	case "[Z":
		return &Object{self, make([]int8, count)}
	case "[B":
		return &Object{self, make([]int8, count)}
	case "[C":
		return &Object{self, make([]uint16, count)}
	case "[S":
		return &Object{self, make([]int16, count)}
	case "[I":
		return &Object{self, make([]int32, count)}
	case "[J":
		return &Object{self, make([]int64, count)}
	case "[F":
		return &Object{self, make([]float32, count)}
	case "[D":
		return &Object{self, make([]float64, count)}
	default:
		return &Object{self, make([]*Object, count)}
	}
}
