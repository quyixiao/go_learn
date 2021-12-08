package heap

import "fmt"
import "go_learn/jvmgo/ch11/classfile"
import "go_learn/jvmgo/ch11/classpath"

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	cp          *classpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}
//loadBasicClasses()函数先加载java.lang.Class类，这又会触发
//java.lang.Object等类和接口的加载。然后遍历classMap，给已经加载 的每一个类关联类对象。
func (self *ClassLoader) loadBasicClasses() {
	jlClassClass := self.LoadClass("java/lang/Class")
	for _, class := range self.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

func (self *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		//loadPrimitiveClasses()方法加载void和基本类型的类
		self.loadPrimitiveClass(primitiveType)
	}
}

func (self *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        className,
		loader:      self,
		initStarted: true,
	}
	class.jClass = self.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	self.classMap[className] = class
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		// already loaded
		return class
	}

	var class *Class
	if name[0] == '[' { // array class ,如果是数组类
		class = self.loadArrayClass(name)
	} else {
		class = self.loadNonArrayClass(name)
	}

	if jlClassClass, ok := self.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}

	return class
}

func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        name,
		loader:      self,
		initStarted: true,						//因为数组类不需要 初始化，所以把initStarted字段设置成true。
		superClass:  self.LoadClass("java/lang/Object"),	//数组类的超类是 java.lang.Object，
		interfaces: []*Class{	//并且实现了java.lang.Cloneable和java.io.Serializable 接口。
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	// 为了确保安全性，Java虚拟机规范要求在执行类的任何代码之前对类进行严格的验证。
	link(class)

	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}

	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5
func (self *ClassLoader) defineClass(data []byte) *Class {
	//方法首先调用parseClass()函数把class文件数据 转换成Class结构体
	//Class结构体的superClass和interfaces字段存放 超类名和直接接口表，这些类名其实都是符号引用
	class := parseClass(data)
	hackClass(class)
	class.loader = self
	//除java.lang.Object以外，所有的类都有且仅有一个 超类。因此，除非是Object类，否则需要递归调用LoadClass()方法 加载它的超类。
	resolveSuperClass(class)
	//resolveInterfaces()函数递归调用 LoadClass()方法加载类的每一个直接接口
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// jvms 5.4.3.1
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// jvms 5.4.2
func prepare(class *Class) {
	//函数计算实例字段的个数
	calcInstanceFieldSlotIds(class)
	//函数计算静态字段的个数
	calcStaticFieldSlotIds(class)
	//函数给类变量分配空间，然后给它们赋 予初始值
	//因为Go语言会保证新创建的Slot结构体有默认值(num字段是 0，ref字段是nil)，而浮点数0编码之后和整数0相同，所以不用做任
	//何操作就可以保证静态变量有默认初始值(数字类型是0，引用类型 是null)。如果静态变量属于基本类型或String类型，有final修饰符，
	//且它的值在编译期已知，则该值存储在class文件常量池中。 initStaticFinalVar()函数从常量池中加载常量值，然后给静态变量赋值
	allocAndInitStaticVars(class)
}

//计算实例变量占据的空间大小
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			//Field结构体的isLongOrDouble()方法返回字段是否是long或 double类型
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}
//初始化静态变量
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {						//如果是静态变量或Final类型的变量，则初始化变量
			initStaticFinalVar(class, field)
		}
	}
}

//因为Go语言会保证新创建的Slot结构体有默认值(num字段是 0，ref字段是nil)，而浮点数0编码之后和整数0相同，所以不用做任 何操作就可以保证
//静态变量有默认初始值(数字类型是0，引用类型 是null)。如果静态变量属于基本类型或String类型，有final修饰符， 且它的值在编译期已知，
//则该值存储在class文件常量池中。 initStaticFinalVar()函数从常量池中加载常量值，然后给静态变量赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}

// todo
func hackClass(class *Class) {
	if class.name == "java/lang/ClassLoader" {
		loadLibrary := class.GetStaticMethod("loadLibrary", "(Ljava/lang/Class;Ljava/lang/String;Z)V")
		loadLibrary.code = []byte{0xb1} // return void
	}
}
