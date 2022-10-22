package main

import "fmt"

func main() {
	cmd := parseCmd()
	//经测试，flag先解析[-option] 再解析args
	println("cmd.class=", cmd.class, "cmd.version=", cmd.versionFlag, "cmd.helpFlag=", cmd.helpFlag)
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Println("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
