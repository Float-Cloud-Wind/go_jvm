package main

import "flag"
import "fmt"
import "os"
// package main: 定义包名，这里是main。
// import "flag": 导入flag包，用于处理命令行参数。
// import "fmt": 导入fmt包，用于格式化输出。
// import "os": 导入os包，用于获取命令行参数。

//定义一个名为Cmd的结构体，用于存储命令行参数和选项。
// java [-options] class [args...]
type Cmd struct {
	//定义一个名为helpFlag的布尔类型字段，表示是否打印帮助信息。
	helpFlag    bool
	//定义一个名为versionFlag的布尔类型字段，表示是否打印版本信息。
	versionFlag bool
	//定义一个名为cpOption的字符串类型字段，表示类路径。
	cpOption    string
	//定义一个名为class的字符串类型字段，表示要执行的类。
	class       string
	//定义一个名为args的切片类型字段，表示类执行时的参数。
	args        []string
}

//解析命令行输入的命令，返回一个Cmd struct 的指针
func parseCmd() *Cmd {
	cmd := &Cmd{}

	//设置flag.Usage函数，用于打印帮助信息
	flag.Usage = printUsage
	//设置helpFlag，用于打印帮助信息，接收四个参数：一个布尔指针、参数的名字、默认值和使用说明
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	//定义一个名为helpFlag的布尔类型字段，并设置其默认值为false，在命令行输入了该选项则为true，表示不打印帮助信息。该字段的别名为"help"和"?"。
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	//定义一个名为versionFlag的布尔类型字段，并设置其默认值为false，表示不打印版本信息。
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	//定义一个名为cpOption的字符串类型字段，并设置其默认值为空字符串，表示没有设置类路径。该字段的别名为"classpath"和"cp"。
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	//设置cpOption的别名，用于设置类路径
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//解析命令行参数，将解析结果存储在Cmd结构体中。
	flag.Parse()

	//获取命令行参数 获取非标志参数
	//非标志参数（non-flag arguments）是指在命令行中出现的不以短横线（-）或双短横线（--）开头的参数
	//myprogram -flag1 value1 arg1 -flag2 value2 arg2 arg3
	//在这个命令行中，-flag1 和 -flag2 是标志参数，arg1、arg2 和 arg3 是非标志参数。
	//-flag1 是标志参数，而 value1 是它的值。同样，-flag2 value2 中的 value2 是 -flag2 的值。
	args := flag.Args()
	//如果参数不为空，则设置class和args
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

//打印帮助信息
func printUsage() {
	//os.Args 是一个字符串切片，它存储了运行程序时提供的所有命令行参数。os.Args[0] 通常是执行的程序名。
	fmt.Printf("Usage: %s [-options] class [args...]\n",  os.Args[0])
	//flag.PrintDefaults()
}