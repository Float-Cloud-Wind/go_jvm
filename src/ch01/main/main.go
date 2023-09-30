package main

import "fmt"

func main() {
	//解析命令行参数，将解析结果封装到一个名为Cmd的结构体中，并返回该结构体的指针
	cmd := parseCmd()

	//是否打印版本信息
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		// 存在help或？选项  或者 没有指定需要执行的java主类,打印帮助信息
		printUsage()
	} else {
		////startJVM，将命令行参数解析后的数据传入
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//打印JVM的参数信息
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}