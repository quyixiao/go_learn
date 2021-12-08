package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

// todo
// go string -> java.lang.String
func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok { //如果 Java字符串已经在池中，直接返回即可
		return internedStr
	}
	//否则先把Go字符串(UTF8 格式)转换成Java字符数组(UTF16格式)，然后创建一个Java字符串 实例，把它的value变量设置成刚刚转换而来的字符数组，最后把
	//Java字符串放入池中。注意，这里其实是跳过了String的构造函数，
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars, nil}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	// 通过反射设置String对象的value值
	jStr.SetRefVar("value", "[C", jChars)	//通过反射赋值

	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

// utf8 -> utf16     先拿到String对象的value变量值，然后把字符数组转换成Go字符串。
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32  Go语言字符串在内存中是UTF8编码的，先把它强制转成 UTF32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16  然后调用utf16包的Encode()函数编码成UTF16
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune  先把UTF16数据转换成UTF8编码，然后强制转换成Go字符串即可。
	return string(runes)
}

// todo
func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	internedStrings[goStr] = jStr
	return jStr
}
