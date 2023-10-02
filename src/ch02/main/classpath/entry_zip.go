package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"
//这段代码定义了一个名为 ZipEntry 的结构体，用于表示一个 ZIP 文件条目，并提供了一些方法来操作这个条目。
//archive/zip（用于读取 ZIP 文件），errors（用于创建错误），io/ioutil（用于读取文件），和 path/filepath（用于处理文件路径）
type ZipEntry struct {
	//用于存储 ZIP 文件的绝对路径
	absPath string
}

//定义了一个名为 newZipEntry 的函数，它接受一个路径作为参数，将这个路径转换为绝对路径，然后创建一个 ZipEntry 实例并返回。
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

//定义了 ZipEntry 结构体的一个方法 readClass，它接受一个类名作为参数，然后在 ZIP 文件中查找这个类，如果找到了，就读取这个类的内容并返回。
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}