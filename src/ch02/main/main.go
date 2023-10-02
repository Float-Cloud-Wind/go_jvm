package main

import "fmt"
import "strings"
import "wsdjvm/main/classpath"

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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:%x\n", classData)
}