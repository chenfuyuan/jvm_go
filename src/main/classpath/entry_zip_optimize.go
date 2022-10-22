package classpath

import (
	"archive/zip"
	"errors"
	"io"
)

type ZipOptimizeEntry struct {
	absPath    string          //绝对路径
	readCloser *zip.ReadCloser //缓存zip压缩包
}

func (z *ZipOptimizeEntry) readClass(className string) ([]byte, Entry, error) {
	//读取文件
	if z.readCloser == nil {
		err := z.openJar()
		if err != nil {
			return nil, nil, err
		}
	}
	//查找文件
	classFile := z.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}
	//读取文件内容
	data, err := readClass(classFile)
	if err != nil {
		return nil, nil, err
	}
	return data, z, err
}

func readClass(file *zip.File) ([]byte, error) {
	readClose, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer readClose.Close()
	data, err := io.ReadAll(readClose)
	if err != nil {
		return nil, err
	}
	return data, err
}

// findClass 查找是否存在文件名为:fileName的文件，找到返回该文件
func (z *ZipOptimizeEntry) findClass(fileName string) *zip.File {
	for _, file := range z.readCloser.File {
		if file.Name == fileName {
			return file
		}
	}
	return nil
}

// openJar 打开jar包
func (z *ZipOptimizeEntry) openJar() error {
	reader, err := zip.OpenReader(z.absPath)
	if err == nil {
		z.readCloser = reader
	}
	return err
}

// String 打印变量
func (z *ZipOptimizeEntry) String() string {
	return z.absPath
}
