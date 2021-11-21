package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// Classpath结构体有三个字段，分别存放三种类路径，Parse()函数使用-Xjre选项解析启动类路径和扩展类路径 。
// 使用-classpath/-cp选项解析用户类路径 。

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

//优先使用用户输入的-Xjre选项作为jre目录，如果没有输入该选项，则在当前目录寻找jre目录，如果找不到，尝试使用
// JAVA_HOME 环境变量，getJreDir()函数的代码如下：
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder !")
}

// 函数用户判断目录是否存在 ，代码如下
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	//启动类加载器来读取
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	//扩展类加载器来读取
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	//用户类加载器来读取文件
	return self.userClasspath.readClass(className)
}

// 注意，传递给ReadClass()方法的类名不包含".class"后缀，最后，String()方法返回用户类路径的字符串表示 ，代码如下
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

