package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool     //帮助命令
	versionFlag bool     //版本命令
	cpOption    string   //用户自定义类路径
	class       string   //执行类名
	args        []string //参数值
}

func parseCmd() *Cmd {
	cmd := &Cmd{} //初始化
	bindVar(cmd)
	getArgs(cmd)
	return cmd
}

func getArgs(cmd *Cmd) {
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
}
func bindVar(cmd *Cmd) {
	flag.Usage = printUsage
	//绑定命令行参数
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")

	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
