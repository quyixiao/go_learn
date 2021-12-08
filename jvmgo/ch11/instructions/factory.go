package instructions

import "fmt"
import "go_learn/jvmgo/ch11/instructions/base"
import . "go_learn/jvmgo/ch11/instructions/comparisons"
import . "go_learn/jvmgo/ch11/instructions/constants"
import . "go_learn/jvmgo/ch11/instructions/control"
import . "go_learn/jvmgo/ch11/instructions/conversions"
import . "go_learn/jvmgo/ch11/instructions/extended"
import . "go_learn/jvmgo/ch11/instructions/loads"
import . "go_learn/jvmgo/ch11/instructions/math"
import . "go_learn/jvmgo/ch11/instructions/references"
import . "go_learn/jvmgo/ch11/instructions/reserved"
import . "go_learn/jvmgo/ch11/instructions/stack"
import . "go_learn/jvmgo/ch11/instructions/stores"

// NoOperandsInstruction singletons
//栈和局部变量操作
//将常量压入栈的指令
//aconst_null 将null对象引用压入栈
//iconst_m1 将int类型常量-1压入栈
//iconst_0 将int类型常量0压入栈
//iconst_1 将int类型常量1压入栈
//iconst_2 将int类型常量2压入栈
//iconst_3 将int类型常量3压入栈
//iconst_4 将int类型常量4压入栈
//iconst_5 将int类型常量5压入栈
//lconst_0 将long类型常量0压入栈
//lconst_1 将long类型常量1压入栈
//fconst_0 将float类型常量0压入栈
//fconst_1 将float类型常量1压入栈
//dconst_0 将double类型常量0压入栈
//dconst_1 将double类型常量1压入栈
//bipush 将一个8位带符号整数压入栈
//sipush 将16位带符号整数压入栈
//ldc 把常量池中的项压入栈 , ldc系列指令把运行时常量池中的常量推到操作数栈顶
//ldc_w 把常量池中的项压入栈（使用宽索引）
//ldc2_w 把常量池中long类型或者double类型的项压入栈（使用宽索引）
//从栈中的局部变量中装载值的指令
//iload 从局部变量中装载int类型值
//lload 从局部变量中装载long类型值
//fload 从局部变量中装载float类型值
//dload 从局部变量中装载double类型值
//aload 从局部变量中装载引用类型值（refernce）
//iload_0 从局部变量0中装载int类型值
//iload_1 从局部变量1中装载int类型值
//iload_2 从局部变量2中装载int类型值
//iload_3 从局部变量3中装载int类型值
//lload_0 从局部变量0中装载long类型值
//lload_1 从局部变量1中装载long类型值
//lload_2 从局部变量2中装载long类型值
//lload_3 从局部变量3中装载long类型值
//fload_0 从局部变量0中装载float类型值
//fload_1 从局部变量1中装载float类型值
//fload_2 从局部变量2中装载float类型值
//fload_3 从局部变量3中装载float类型值
//dload_0 从局部变量0中装载double类型值
//dload_1 从局部变量1中装载double类型值
//dload_2 从局部变量2中装载double类型值
//dload_3 从局部变量3中装载double类型值
//aload_0 从局部变量0中装载引用类型值
//aload_1 从局部变量1中装载引用类型值
//aload_2 从局部变量2中装载引用类型值
//aload_3 从局部变量3中装载引用类型值
//iaload 从数组中装载int类型值
//laload 从数组中装载long类型值
//faload 从数组中装载float类型值
//daload 从数组中装载double类型值
//aaload 从数组中装载引用类型值
//baload 从数组中装载byte类型或boolean类型值
//caload 从数组中装载char类型值
//saload 从数组中装载short类型值
//将栈中的值存入局部变量的指令
//istore 将int类型值存入局部变量
//lstore 将long类型值存入局部变量
//fstore 将float类型值存入局部变量
//dstore 将double类型值存入局部变量
//astore 将将引用类型或returnAddress类型值存入局部变量
//istore_0 将int类型值存入局部变量0
//istore_1 将int类型值存入局部变量1
//istore_2 将int类型值存入局部变量2
//istore_3 将int类型值存入局部变量3
//lstore_0 将long类型值存入局部变量0
//lstore_1 将long类型值存入局部变量1
//lstore_2 将long类型值存入局部变量2
//lstore_3 将long类型值存入局部变量3
//fstore_0 将float类型值存入局部变量0
//fstore_1 将float类型值存入局部变量1
//fstore_2 将float类型值存入局部变量2
//fstore_3 将float类型值存入局部变量3
//dstore_0 将double类型值存入局部变量0
//dstore_1 将double类型值存入局部变量1
//dstore_2 将double类型值存入局部变量2
//dstore_3 将double类型值存入局部变量3
//astore_0 将引用类型或returnAddress类型值存入局部变量0
//astore_1 将引用类型或returnAddress类型值存入局部变量1
//astore_2 将引用类型或returnAddress类型值存入局部变量2
//astore_3 将引用类型或returnAddress类型值存入局部变量3
//iastore 将int类型值存入数组中
//lastore 将long类型值存入数组中
//fastore 将float类型值存入数组中
//dastore 将double类型值存入数组中
//aastore 将引用类型值存入数组中
//bastore 将byte类型或者boolean类型值存入数组中
//castore 将char类型值存入数组中
//sastore 将short类型值存入数组中
//wide指令
//wide 使用附加字节扩展局部变量索引
//通用(无类型）栈操作
//nop 不做任何操作
//pop 弹出栈顶端一个字长的内容
//pop2 弹出栈顶端两个字长的内容
//dup 复制栈顶部一个字长内容
//dup_x1 复制栈顶部一个字长的内容，然后将复制内容及原来弹出的两个字长的内容压入栈
//
//dup_x2 复制栈顶部一个字长的内容，然后将复制内容及原来弹出的三个字长的内容压入栈
//dup2 复制栈顶部两个字长内容
//dup2_x1 复制栈顶部两个字长的内容，然后将复制内容及原来弹出的三个字长的内容压入栈
//dup2_x2 复制栈顶部两个字长的内容，然后将复制内容及原来弹出的四个字长的内容压入栈
//swap 交换栈顶部两个字长内容
//类型转换
//i2l 把int类型的数据转化为long类型
//i2f 把int类型的数据转化为float类型
//i2d 把int类型的数据转化为double类型
//l2i 把long类型的数据转化为int类型
//l2f 把long类型的数据转化为float类型
//l2d 把long类型的数据转化为double类型
//f2i 把float类型的数据转化为int类型
//f2l 把float类型的数据转化为long类型
//f2d 把float类型的数据转化为double类型
//d2i 把double类型的数据转化为int类型
//d2l 把double类型的数据转化为long类型
//d2f 把double类型的数据转化为float类型
//i2b 把int类型的数据转化为byte类型
//i2c 把int类型的数据转化为char类型
//i2s 把int类型的数据转化为short类型
//整数运算
//iadd 执行int类型的加法
//ladd 执行long类型的加法
//isub 执行int类型的减法
//lsub 执行long类型的减法
//imul 执行int类型的乘法
//lmul 执行long类型的乘法
//idiv 执行int类型的除法
//ldiv 执行long类型的除法
//irem 计算int类型除法的余数
//lrem 计算long类型除法的余数
//ineg 对一个int类型值进行取反操作
//lneg 对一个long类型值进行取反操作
//iinc 把一个常量值加到一个int类型的局部变量上
//逻辑运算
//移位操作
//ishl 执行int类型的向左移位操作
//lshl 执行long类型的向左移位操作
//ishr 执行int类型的向右移位操作
//lshr 执行long类型的向右移位操作
//iushr 执行int类型的向右逻辑移位操作
//lushr 执行long类型的向右逻辑移位操作
//按位布尔运算
//iand 对int类型值进行“逻辑与”操作
//land 对long类型值进行“逻辑与”操作
//ior 对int类型值进行“逻辑或”操作
//lor 对long类型值进行“逻辑或”操作
//ixor 对int类型值进行“逻辑异或”操作
//lxor 对long类型值进行“逻辑异或”操作
//浮点运算
//fadd 执行float类型的加法
//dadd 执行double类型的加法
//fsub 执行float类型的减法
//dsub 执行double类型的减法
//fmul 执行float类型的乘法
//dmul 执行double类型的乘法
//fdiv 执行float类型的除法
//ddiv 执行double类型的除法
//frem 计算float类型除法的余数
//drem 计算double类型除法的余数
//fneg 将一个float类型的数值取反
//dneg 将一个double类型的数值取反
//对象和数组
//对象操作指令
//new 创建一个新对象
//checkcast 确定对象为所给定的类型
//getfield 从对象中获取字段
//putfield [putfield和getfield用 于存取实例变量;]
//getstatic
//putstatic  [putstatic和getstatic指令用于存取静态变量]
//instanceof [指令用于判断对象是否属于 某种类型;]
//数组操作指令
//newarray 分配数据成员类型为基本上数据类型的新数组
//anewarray 分配数据成员类型为引用类型的新数组
//arraylength 获取数组长度
//multianewarray 分配新的多维数组
//控制流
//条件分支指令
//ifeq 如果等于0，则跳转
//ifne 如果不等于0，则跳转
//iflt 如果小于0，则跳转
//ifge 如果大于等于0，则跳转
//ifgt 如果大于0，则跳转
//ifle 如果小于等于0，则跳转
//if_icmpcq 如果两个int值相等，则跳转
//if_icmpne 如果两个int类型值不相等，则跳转
//if_icmplt 如果一个int类型值小于另外一个int类型值，则跳转
//if_icmpge 如果一个int类型值大于或者等于另外一个int类型值，则跳转
//if_icmpgt 如果一个int类型值大于另外一个int类型值，则跳转
//if_icmple 如果一个int类型值小于或者等于另外一个int类型值，则跳转
//ifnull 如果等于null，则跳转
//ifnonnull 如果不等于null，则跳转
//if_acmpeq 如果两个对象引用相等，则跳转
//if_acmpnc 如果两个对象引用不相等，则跳转
//比较指令
//lcmp 比较long类型值
//fcmpl 比较float类型值（当遇到NaN时，返回-1）
//fcmpg 比较float类型值（当遇到NaN时，返回1）
//dcmpl 比较double类型值（当遇到NaN时，返回-1）
//dcmpg 比较double类型值（当遇到NaN时，返回1）
//无条件转移指令
//goto 无条件跳转
//goto_w 无条件跳转（宽索引）
//表跳转指令
//tableswitch 通过索引访问跳转表，并跳转
//lookupswitch 通过键值匹配访问跳转表，并执行跳转操作
//异常
//athrow 抛出异常或错误
//finally子句
//jsr 跳转到子例程
//jsr_w 跳转到子例程（宽索引）
//rct 从子例程返回
//方法调用与返回
//方法调用指令
//invokcvirtual 运行时按照对象的类来调用实例方法
//invokespecial 根据编译时类型来调用实例方法
//invokestatic 调用类（静态）方法
//invokcinterface 调用接口方法
//方法返回指令
//ireturn 从方法中返回int类型的数据
//lreturn 从方法中返回long类型的数据
//freturn 从方法中返回float类型的数据
//dreturn 从方法中返回double类型的数据
//areturn 从方法中返回引用类型的数据
//return 从方法中返回，返回值为void
//线程同步
//montiorenter 进入并获取对象监视器
//monitorexit 释放并退出对象监视器
//***********************************************************************
//JVM指令助记符
//变量到操作数栈：iload,iload_,lload,lload_,fload,fload_,dload,dload_,aload,aload_
//操作数栈到变量：istore,istore_,lstore,lstore_,fstore,fstore_,dstore,dstor_,astore,astore_
//常数到操作数栈：bipush,sipush,ldc,ldc_w,ldc2_w,aconst_null,iconst_ml,iconst_,lconst_,fconst_,dconst_
//加：iadd,ladd,fadd,dadd
//减：isub,lsub,fsub,dsub
//乘：imul,lmul,fmul,dmul
//除：idiv,ldiv,fdiv,ddiv
//余数：irem,lrem,frem,drem
//取负：ineg,lneg,fneg,dneg
//移位：ishl,lshr,iushr,lshl,lshr,lushr
//按位或：ior,lor
//按位与：iand,land
//按位异或：ixor,lxor
//类型转换：i2l,i2f,i2d,l2f,l2d,f2d(放宽数值转换)
//i2b,i2c,i2s,l2i,f2i,f2l,d2i,d2l,d2f(缩窄数值转换)
//创建类实便：new
//创建新数组：newarray,anewarray,multianwarray
//访问类的域和类实例域：getfield,putfield,getstatic,putstatic
//把数据装载到操作数栈：baload,caload,saload,iaload,laload,faload,daload,aaload
//从操作数栈存存储到数组：bastore,castore,sastore,iastore,lastore,fastore,dastore,aastore
//获取数组长度：arraylength
//检相类实例或数组属性：instanceof,checkcast
//操作数栈管理：pop,pop2,dup,dup2,dup_xl,dup2_xl,dup_x2,dup2_x2,swap
//有条件转移：ifeq,iflt,ifle,ifne,ifgt,ifge,ifnull,ifnonnull,if_icmpeq,if_icmpene,
//if_icmplt,if_icmpgt,if_icmple,if_icmpge,if_acmpeq,if_acmpne,lcmp,fcmpl
//fcmpg,dcmpl,dcmpg
//复合条件转移：tableswitch,lookupswitch
//无条件转移：goto,goto_w,jsr,jsr_w,ret
//调度对象的实便方法：invokevirtual
//调用由接口实现的方法：invokeinterface
//调用需要特殊处理的实例方法：invokespecial
//调用命名类中的静态方法：invokestatic
//方法返回：ireturn,lreturn,freturn,dreturn,areturn,return
//异常：athrow
//finally关键字的实现使用：jsr,jsr_w,ret



var (
	nop           = &NOP{}
	aconst_null   = &ACONST_NULL{}
	iconst_m1     = &ICONST_M1{}
	iconst_0      = &ICONST_0{}
	iconst_1      = &ICONST_1{}
	iconst_2      = &ICONST_2{}
	iconst_3      = &ICONST_3{}
	iconst_4      = &ICONST_4{}
	iconst_5      = &ICONST_5{}
	lconst_0      = &LCONST_0{}
	lconst_1      = &LCONST_1{}
	fconst_0      = &FCONST_0{}
	fconst_1      = &FCONST_1{}
	fconst_2      = &FCONST_2{}
	dconst_0      = &DCONST_0{}
	dconst_1      = &DCONST_1{}
	iload_0       = &ILOAD_0{}
	iload_1       = &ILOAD_1{}
	iload_2       = &ILOAD_2{}
	iload_3       = &ILOAD_3{}
	lload_0       = &LLOAD_0{}
	lload_1       = &LLOAD_1{}
	lload_2       = &LLOAD_2{}
	lload_3       = &LLOAD_3{}
	fload_0       = &FLOAD_0{}
	fload_1       = &FLOAD_1{}
	fload_2       = &FLOAD_2{}
	fload_3       = &FLOAD_3{}
	dload_0       = &DLOAD_0{}
	dload_1       = &DLOAD_1{}
	dload_2       = &DLOAD_2{}
	dload_3       = &DLOAD_3{}
	aload_0       = &ALOAD_0{}
	aload_1       = &ALOAD_1{}
	aload_2       = &ALOAD_2{}
	aload_3       = &ALOAD_3{}
	iaload        = &IALOAD{}
	laload        = &LALOAD{}
	faload        = &FALOAD{}
	daload        = &DALOAD{}
	aaload        = &AALOAD{}
	baload        = &BALOAD{}
	caload        = &CALOAD{}
	saload        = &SALOAD{}
	istore_0      = &ISTORE_0{}
	istore_1      = &ISTORE_1{}
	istore_2      = &ISTORE_2{}
	istore_3      = &ISTORE_3{}
	lstore_0      = &LSTORE_0{}
	lstore_1      = &LSTORE_1{}
	lstore_2      = &LSTORE_2{}
	lstore_3      = &LSTORE_3{}
	fstore_0      = &FSTORE_0{}
	fstore_1      = &FSTORE_1{}
	fstore_2      = &FSTORE_2{}
	fstore_3      = &FSTORE_3{}
	dstore_0      = &DSTORE_0{}
	dstore_1      = &DSTORE_1{}
	dstore_2      = &DSTORE_2{}
	dstore_3      = &DSTORE_3{}
	astore_0      = &ASTORE_0{}
	astore_1      = &ASTORE_1{}
	astore_2      = &ASTORE_2{}
	astore_3      = &ASTORE_3{}
	iastore       = &IASTORE{}
	lastore       = &LASTORE{}
	fastore       = &FASTORE{}
	dastore       = &DASTORE{}
	aastore       = &AASTORE{}
	bastore       = &BASTORE{}
	castore       = &CASTORE{}
	sastore       = &SASTORE{}
	pop           = &POP{}
	pop2          = &POP2{}
	dup           = &DUP{}
	dup_x1        = &DUP_X1{}
	dup_x2        = &DUP_X2{}
	dup2          = &DUP2{}
	dup2_x1       = &DUP2_X1{}
	dup2_x2       = &DUP2_X2{}
	swap          = &SWAP{}
	iadd          = &IADD{}
	ladd          = &LADD{}
	fadd          = &FADD{}
	dadd          = &DADD{}
	isub          = &ISUB{}
	lsub          = &LSUB{}
	fsub          = &FSUB{}
	dsub          = &DSUB{}
	imul          = &IMUL{}
	lmul          = &LMUL{}
	fmul          = &FMUL{}
	dmul          = &DMUL{}
	idiv          = &IDIV{}
	ldiv          = &LDIV{}
	fdiv          = &FDIV{}
	ddiv          = &DDIV{}
	irem          = &IREM{}
	lrem          = &LREM{}
	frem          = &FREM{}
	drem          = &DREM{}
	ineg          = &INEG{}
	lneg          = &LNEG{}
	fneg          = &FNEG{}
	dneg          = &DNEG{}
	ishl          = &ISHL{}
	lshl          = &LSHL{}
	ishr          = &ISHR{}
	lshr          = &LSHR{}
	iushr         = &IUSHR{}
	lushr         = &LUSHR{}
	iand          = &IAND{}
	land          = &LAND{}
	ior           = &IOR{}
	lor           = &LOR{}
	ixor          = &IXOR{}
	lxor          = &LXOR{}
	i2l           = &I2L{}
	i2f           = &I2F{}
	i2d           = &I2D{}
	l2i           = &L2I{}
	l2f           = &L2F{}
	l2d           = &L2D{}
	f2i           = &F2I{}
	f2l           = &F2L{}
	f2d           = &F2D{}
	d2i           = &D2I{}
	d2l           = &D2L{}
	d2f           = &D2F{}
	i2b           = &I2B{}
	i2c           = &I2C{}
	i2s           = &I2S{}
	lcmp          = &LCMP{}
	fcmpl         = &FCMPL{}
	fcmpg         = &FCMPG{}
	dcmpl         = &DCMPL{}
	dcmpg         = &DCMPG{}
	ireturn       = &IRETURN{}
	lreturn       = &LRETURN{}
	freturn       = &FRETURN{}
	dreturn       = &DRETURN{}
	areturn       = &ARETURN{}
	_return       = &RETURN{}
	arraylength   = &ARRAY_LENGTH{}
	athrow        = &ATHROW{}
	monitorenter  = &MONITOR_ENTER{}
	monitorexit   = &MONITOR_EXIT{}
	invoke_native = &INVOKE_NATIVE{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		// 例如指令iconst_<i> 可以压入int常量-1，0，1，2，3，4，5 ,const_0 表示把int类型的0值压入操作数栈
		// 这样iconst_0就不需要专门为入栈操作数直接操作数值了，而且也避免了操作数的读取和解析步骤，在本例子中，把压入
		// 0这个操作的指令由iconst_0改为bipush0 也能获取正确的结果，但是编译码会因此额外增加了1个字节的长度，简单实现
		// 的虚拟机可能要在每次循环时消耗更多的时间用于获取和解析这个操作数，因此使用隐式的操作数可以让编译后的代码更加简洁，更加高效
		return iconst_0
	case 0x04:
		return iconst_1			// 将1压入操作数栈
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:											// 16 * 1 + 0  = 16
		return &BIPUSH{}
	case 0x11:
		return &SIPUSH{}
		// ldc系列指令从运行时常量池中加载常量值，并把它推入操作 数栈。ldc系列指令属于常量类指令，共3条。其中ldc和ldc_w指令用 于加载int、float和字符串常量
	case 0x12:
		return &LDC{}
	case 0x13:
		return &LDC_W{}
	case 0x14:
		return &LDC2_W{}
	case 0x15:
		return &ILOAD{}
	case 0x16:
		return &LLOAD{}
	case 0x17:
		return &FLOAD{}
	case 0x18:
		return &DLOAD{}
	case 0x19:
		return &ALOAD{}
	case 0x1a:
		return iload_0
	case 0x1b:
		// iload_1指令的作用是将第一个局部变量压入到操作数栈  16 + 11 = 27
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	case 0x2e:
		return iaload
	case 0x2f:
		return laload
	case 0x30:
		return faload
	case 0x31:
		return daload
	case 0x32:
		return aaload
	case 0x33:
		return baload
	case 0x34:
		return caload
	case 0x35:
		return saload
	case 0x36:
		return &ISTORE{}
	case 0x37:
		return &LSTORE{}
	case 0x38:
		return &FSTORE{}
	case 0x39:
		return &DSTORE{}
	case 0x3a:
		return &ASTORE{}
	case 0x3b:
		return istore_0
	case 0x3c:														// 3 * 16 + 12 = 60
		// istore_1 指令的作用是从操作数栈中弹出一个int类型的值，并保存在第一个局部变量中
		return istore_1
	case 0x3d:									// 3 * 16 + 13 = 61
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	case 0x4f:
		return iastore
	case 0x50:
		return lastore
	case 0x51:
		return fastore
	case 0x52:
		return dastore
	case 0x53:
		return aastore
	case 0x54:
		return bastore
	case 0x55:
		return castore
	case 0x56:
		return sastore
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return &IINC{}
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
		//条件分支
	case 0x99:
		return &IFEQ{}
	case 0x9a:
		return &IFNE{}
	case 0x9b:
		return &IFLT{}
	case 0x9c:
		return &IFGE{}
	case 0x9d:
		return &IFGT{}
	case 0x9e:
		return &IFLE{}
	case 0x9f:
		return &IF_ICMPEQ{}
	case 0xa0:
		return &IF_ICMPNE{}
	case 0xa1:
		return &IF_ICMPLT{}
	case 0xa2:
		return &IF_ICMPGE{}
	case 0xa3:
		return &IF_ICMPGT{}
	case 0xa4:
		return &IF_ICMPLE{}
	case 0xa5:
		return &IF_ACMPEQ{}
	case 0xa6:
		return &IF_ACMPNE{}
	case 0xa7:
		return &GOTO{}
	// case 0xa8:
	// 	return &JSR{}
	// case 0xa9:
	// 	return &RET{}
	// 复合条件分支
	case 0xaa:
		return &TABLE_SWITCH{}
	case 0xab:
		return &LOOKUP_SWITCH{}
	case 0xac:
		return ireturn
	case 0xad:
		return lreturn
	case 0xae:
		return freturn
	case 0xaf:
		return dreturn
	case 0xb0:
		return areturn
	case 0xb1:
		return _return
	case 0xb2:
		return &GET_STATIC{}
	case 0xb3:
		return &PUT_STATIC{}
	case 0xb4:
		return &GET_FIELD{}
	case 0xb5:
		return &PUT_FIELD{}
	case 0xb6:
		return &INVOKE_VIRTUAL{}	//
	case 0xb7:
		return &INVOKE_SPECIAL{}
	case 0xb8:
		return &INVOKE_STATIC{}
	case 0xb9:
		return &INVOKE_INTERFACE{}
	// case 0xba:
	// 	return &INVOKE_DYNAMIC{}
	case 0xbb:
		return &NEW{}
	case 0xbc:
		return &NEW_ARRAY{}
	case 0xbd:
		return &ANEW_ARRAY{}
	case 0xbe:
		return arraylength
	case 0xbf:
		return athrow			//
	case 0xc0:
		return &CHECK_CAST{}
	case 0xc1:
		return &INSTANCE_OF{}
	case 0xc2:
		return monitorenter
	case 0xc3:
		return monitorexit
	case 0xc4:
		return &WIDE{}
	case 0xc5:
		return &MULTI_ANEW_ARRAY{}
	case 0xc6:
		return &IFNULL{}
	case 0xc7:
		return &IFNONNULL{}
		//无条件分支：goto ,goto_w ,jsr,jsr_w,ret
	case 0xc8:
		return &GOTO_W{}
	// case 0xc9:
	// 	return &JSR_W{}
	// case 0xca: breakpoint
	case 0xfe:
		return invoke_native
	// case 0xff: impdep2
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
