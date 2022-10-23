package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string //绝对路径
}

func newDirEntry(path string) *DirEntry {
	//转换绝对路径
	absDir := loadAbsPath(path)
	//创建对象
	return &DirEntry{absDir}
}

func (d *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	//拼接生成文件的绝对路径
	fileName := filepath.Join(d.absDir, className)
	//读取文件
	data, err := os.ReadFile(fileName)
	return data, d, err
}

func (d *DirEntry) String() string {
	return d.absDir
}
