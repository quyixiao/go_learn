package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
//把类名转变成类型描述符，然后在前面加上方括号即可。
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
//数组类名以方括号开头，把它去掉就是数组元素的类型描述
//符，然后把类型描述符转成类名即可。
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

// [XXX => [XXX
// int  => I
// XXX  => LXXX;
//果是数组类名，描述符就是类名，直接返回即可。如果是基 本类型名，返回对应的类型描述符，否则肯定是普通的类名，
//前面 加上方括号，结尾加上分号即可得到类型描述符
func toDescriptor(className string) string {
	if className[0] == '[' {
		// array
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		// primitive
		return d
	}
	// object
	return "L" + className + ";"
}

// [XXX  => [XXX
// LXXX; => XXX
// I     => int

func toClassName(descriptor string) string {
	//如果类型描述符以方括号开头，那么肯定是数组，描述符即是 类名。
	if descriptor[0] == '[' {
		// array
		return descriptor
	}
	//如果类型描述符以L开头，那么肯定是类描述符，去掉开头的L和末尾的分号即是类名，否则判断是否是基本类型的描述符
	if descriptor[0] == 'L' {
		// object
		return descriptor[1 : len(descriptor)-1]
	}
	//如果是，返回基本类型名称，
	for className, d := range primitiveTypes {
		if d == descriptor {
			// primitive
			return className
		}
	}
	//否则调用panic()函数终止程序执行。
	panic("Invalid descriptor: " + descriptor)
}
