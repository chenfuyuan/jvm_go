package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// PathListSeparator 系统分隔符
const PathListSeparator = string(os.PathListSeparator)

type Entry interface {
	//readClass 读取class文件
	//@param className class文件路径
	//@return []byte 读取到的字节数据
	//@return Entry 定位到class的Entry
	//@return error 错误信息
	///**
	readClass(className string) ([]byte, Entry, error)

	// String /**
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, PathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return nil
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

func loadAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return absPath
}
