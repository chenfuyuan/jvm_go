package main

import (
	"fmt"
	"go_jvm/src/main/classpath"
	"os"
	"strings"
)

func main() {
	dir, _ := os.UserHomeDir()
	fmt.Printf("userHome:%v\n", dir)
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("jreOption:%v cpOption:%v\n", cmd.xJreOption, cmd.cpOption)
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:%v\n", classData)
}
