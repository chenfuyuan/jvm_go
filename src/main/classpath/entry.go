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
		return newWildcardEntry(path)
	}
	if judgeIsZip(path) {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

func judgeIsZip(fileName string) bool {
	return judgeIsJar(fileName) ||
		strings.HasSuffix(fileName, ".zip") ||
		strings.HasSuffix(fileName, ".ZIP")
}

func judgeIsJar(fileName string) bool {
	return strings.HasSuffix(fileName, ".jar") ||
		strings.HasSuffix(fileName, ".JAR")
}

func loadAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return absPath
}
