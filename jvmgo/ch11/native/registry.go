package native

import "go_learn/jvmgo/ch11/rtda"

//把本地方法定义成一个函数，参数是Frame结构体指针
//这个frame参数就是本地方法的工作空间，也就是连接Java 虚拟机和Java类库的桥梁，后面会看到它是如何发挥作用的。
type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	//类名、方法名和方法描述符加在一起才能唯一确定一个方法， 所以把它们的组合作为本地方法注册表的键
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

//FindNativeMethod()方法根据类名、方法名和方法描述符查找 本地方法实现，如果找不到，则返回nil
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" {
		if methodName == "registerNatives" || methodName == "initIDs" {
			return emptyNativeMethod
		}
	}
	return nil
}
