package heap

type Object struct {
	class *Class      //一个存放对象的Class指针
	data  interface{} // Slots for Object, []int32 for int[] ...
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}




// reflection
//先拿到String对象的value变量值，然后把字符数组转换成Go字 符串。
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

//Go语言字符串在内存中是UTF8编码的，先把它强制转成 UTF32，然后调用utf16包的Encode()函数编码成UTF16。Object结 构体的SetRefVar()
//方法直接给对象的引用类型实例变量赋值
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
