package main

func main() {
	//cmd := parseCmd()



	cmd := &Cmd{
		helpFlag:    false,
		versionFlag: false,
		cpOption:    "/Users/quyixiao/gitlab/jvmdemo/target/classes/jvmgo/book/ch06",
		class:       "MyObject",
		XjreOption:  "/Library/Java/JavaVirtualMachines/jdk1.8.0_131.jdk/Contents/Home/jre",
		args:        []string{},
		verboseClassFlag: true,
	}
	if cmd.versionFlag {
		println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		newJVM(cmd).start()
	}
}
