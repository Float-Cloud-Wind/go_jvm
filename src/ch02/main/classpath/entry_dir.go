package classpath

import "io/ioutil"
import "path/filepath"
//import "io/ioutil": 导入io/ioutil包，用于读取文件。
//import "path/filepath": 导入path/filepath包，用于处理文件路径。
type DirEntry struct {
	//存放目录的绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	//使用filepath.Abs函数将给定的路径转换为绝对路径，并将结果存储在absDir变量中。同时，将错误信息存储在err变量中。
	//filepath.Abs("./test") 将相对路径 ./test 转换为绝对路径。例如，如果你在 /home/user 目录下运行这个程序，那么输出的结果将是 /home/user/test。
	absDir, err := filepath.Abs(path)
	//panic 是一个内置函数，用于生成一个运行时错误，并停止当前的 Go 程序。当你调用 panic 函数时，程序的正常执行流程会立即停止，所有的延迟函数（deferred functions）会被执行，然后程序会返回到 Go 运行时系统。在你给出的代码中，panic(err) 被用于处理 filepath.Abs(path) 函数可能产生的错误。如果 filepath.Abs(path) 函数返回一个非 nil 的错误，那么 panic(err) 会立即停止程序的执行，并打印出错误信息。
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

//先把目录和class文件名拼成一个完整的路径，然后调用ioutil包提供的ReadFile（）函数读取class文件内容，最后返回
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	//filepath.Join 是 Go 语言 filepath 包中的一个函数，它的作用是将一个或多个路径片段连接成一个单一的路径。这个函数会清理路径，解析相对路径和符号链接，并返回结果路径的规范化形式。
	//在你给出的代码中，filepath.Join(self.absDir, className) 将 self.absDir（一个目录的绝对路径）和 className（一个类名）连接成一个完整的文件路径。这个文件路径可以用于读取或写入类文件。
	//例如，如果 self.absDir 是 /path/to/dir，className 是 MyClass.class，那么 filepath.Join(self.absDir, className) 的结果将是 /path/to/dir/MyClass.class。
	fileName := filepath.Join(self.absDir, className)
	//ioutil.ReadFile 是 Go 语言 io/ioutil 包中的一个函数，它的作用是读取指定文件的全部内容，并返回一个字节切片（byte slice）。如果在读取文件过程中发生错误，这个函数会返回一个非 nil 的错误。
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}