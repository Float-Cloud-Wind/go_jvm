package classpath

import "os"
import "strings"

// :(linux/unix) or ;(windows)
//代表了系统特定的路径列表分隔符。在 Unix 系统（如 Linux 或 Mac OS）中，这个分隔符是冒号 :，而在 Windows 系统中，这个分隔符是分号 ;
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	//readClass（）方法负责寻找和加载class文件
	// className: fully/qualified/ClassName.class
	readClass(className string) ([]byte, Entry, error)
	//返回变量的字符串表示
	String() string
}

//根据参数创建不同类型的Entry实例
//Entry接口有4个实现
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	//path 以 * 结尾
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

	//表示目录形式的类路径
	return newDirEntry(path)
}