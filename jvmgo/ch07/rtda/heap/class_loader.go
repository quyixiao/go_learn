package heap

import "fmt"
import "go_learn/jvmgo/ch07/classfile"
import "go_learn/jvmgo/ch07/classpath"

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	cp       *classpath.Classpath // ClassLoader依赖Classpath来搜索和读取class文件，cp字段保存 Classpath指针
	classMap map[string]*Class    // loaded classes	//classMap字段记录已经加载的类数据，key是类的完 全限定名。
	//方法区一直只是个抽象的概念，现在可 以把classMap字段当作方法区的具体实现。
}

//NewClassLoader()函数 创建ClassLoader实例
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

//首先找到class文 件并把数据读取到内存;然后解析class文件，生成虚拟机可以使用的类数据，并放入方法区;最后进行链接。
//LoadClass()方法把类数据加载到方法区
//先查找classMap，看类是否已经被加载。如果是，直接返回类 数据，否则调用loadNonArrayClass()方法加载类。数组类和普通类 有很大的不同，它的数据并不是来自class文件，而是由Java虚拟机 在运行期间生成。
func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		// already loaded
		return class
	}

	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

//readClass()方法只是调用了Classpath的ReadClass()方法，并进 行了错误处理。需要解释一下它的返回值。为了打印类加载信息， 把最终加载class文件的类路径项也返回给了调用者。
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	fmt.Println(" className ======" + name)
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5
//defineClass()方法首先调用parseClass()函数把class文件数据 转换成Class结构体。
//Class结构体的superClass和interfaces字段存放 超类名和直接接口表，这些类名其实都是符号引用。
//根据Java虚拟 机规范的5.3.5节，调用resolveSuperClass()和resolveInterfaces()函数 解析这些类符号引用。
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
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
//除java.lang.Object以外，所有的类都有且仅有一个 超类。因此，除非是Object类，否则需要递归调用LoadClass()方法 加载它的超类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

//resolveInterfaces()函数递归调用 LoadClass()方法加载类的每一个直接接口
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

//类的链接分为验证和准备两个必要阶段，link()方法的代码
func link(class *Class) {
	//为了确保安全性，Java虚拟机规范要求在执行类的任何代码之 前，对类进行严格的验证。
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// jvms 5.4.2
func prepare(class *Class) {
	//calcInstanceFieldSlotIds()函数计算实例字段的个数，同时给它 们编号
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

//第一个问题比较好解决，只要数一下类的字段即可。假设某个 类有m个静态字段和n个实例字段，那么静态变量和实例变量所需 的空间大小就分别是m'和n'。
//这里要注意两点。首先，类是可以继承 的。也就是说，在数实例变量时，要递归地数超类的实例变量;其 次，long和double字段都占据两个位置，所以m'>=m，n'>=n。
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

//第二个问题也不算难，在数字段时，给字段按顺序编上号就可 以了。这里有三点需要要注意。首先，静态字段和实例字段要分开 编号，否则会混乱。
//其次，对于实例字段，一定要从继承关系的最 顶端，也就是java.lang.Object开始编号，否则也会混乱。最后，编号 时也要考虑long和double类型。
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			//Field结构体的isLongOrDouble()方法返回字段是否是long或 double类型
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

//allocAndInitStaticVars()函数给类变量分配空间，然后给它们赋 予初始值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

//因为Go语言会保证新创建的Slot结构体有默认值(num字段是 0，ref字段是nil)，而浮点数0编码之后和整数0相同，所以不用做任 何操作就可以保证静
//态变量有默认初始值(数字类型是0，引用类型 是null)。如果静态变量属于基本类型或String类型，有final修饰符， 且它的值在编译期已知，
//则该值存储在class文件常量池中。 initStaticFinalVar()函数从常量池中加载常量值，然后给静态变量 赋值
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
			panic("todo")
		}
	}
}
