package lang

import "runtime"
import "time"
import "go_learn/jvmgo/ch11/instructions/base"
import "go_learn/jvmgo/ch11/native"
import "go_learn/jvmgo/ch11/rtda"
import "go_learn/jvmgo/ch11/rtda/heap"

const jlSystem = "java/lang/System"

func init() {
	native.Register(jlSystem, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register(jlSystem, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	native.Register(jlSystem, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	native.Register(jlSystem, "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	native.Register(jlSystem, "setErr0", "(Ljava/io/PrintStream;)V", setErr0)
	native.Register(jlSystem, "currentTimeMillis", "()J", currentTimeMillis)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
//先从局部变量表中拿到5个参数。源数组和目标数组都不能是 null，否则按照System类的Javadoc应该抛出NullPointerException异常
func arraycopy(frame *rtda.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	//源数组和目标数组必须兼容才能拷贝，否则应该抛出 ArrayStoreException异常
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	//checkArrayCopy()函数的代码稍后给出。接下来检查srcPos、 destPos和length参数，如果有问题则抛出 IndexOutOfBoundsException异常
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	//最后，参数合法，调用ArrayCopy()函数进行数组拷贝
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

//checkArrayCopy()函数首先确保src和dest都是数组，然后检查 数组类型。如果两者都是引用数组，则可以拷贝，否则两者必须是 相同类型的基本类型数组
func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	//Class结构体的IsPrimitive()函数判断类是否是基本类型的类
	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
	vars := frame.LocalVars()
	props := vars.GetRef(0)

	stack := frame.OperandStack()
	stack.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread()
	for key, val := range _sysProps() {
		jKey := heap.JString(frame.Method().Class().Loader(), key)
		jVal := heap.JString(frame.Method().Class().Loader(), val)
		ops := rtda.NewOperandStack(3)
		ops.PushRef(props)
		ops.PushRef(jKey)
		ops.PushRef(jVal)
		shimFrame := rtda.NewShimFrame(thread, ops)
		thread.PushFrame(shimFrame)

		base.InvokeMethod(shimFrame, setPropMethod)
	}
}

func _sysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https://github.com/zxh0/jvm.go",
		"java.home":            "todo",
		"java.class.version":   "52.0",
		"java.class.path":      "todo",
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	in := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("in", "Ljava/io/InputStream;", in)
}

// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	out := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("out", "Ljava/io/PrintStream;", out)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	err := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("err", "Ljava/io/PrintStream;", err)
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *rtda.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	stack := frame.OperandStack()
	stack.PushLong(millis)
}
