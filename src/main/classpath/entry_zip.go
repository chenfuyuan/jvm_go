// Package classpath class文件加载
package classpath

import (
	"archive/zip"
	"errors"
	"io"
)

type ZipEntry struct {
	absPath string //绝对路径
}

// newZipEntry 构造函数
func newZipEntry(path string) *ZipEntry {
	absPath := loadAbsPath(path)
	return &ZipEntry{absPath}
}

func (z *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(z.absPath) //打开读流
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()               //使用完后关闭reader
	for _, file := range reader.File { //遍历其中的文件
		if file.Name == className { //如果文件名称等于所要执行的类
			readCloser, err := file.Open() //打开文件
			if err != nil {
				return nil, nil, err
			}
			defer readCloser.Close()            //延迟关闭流
			data, err := io.ReadAll(readCloser) //读取所有数据
			if err != nil {
				return nil, nil, err
			}
			return data, z, nil
		}
	}
	//class未找到
	return nil, nil, errors.New("class not found: " + className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}
