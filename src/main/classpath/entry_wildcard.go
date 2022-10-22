package classpath

import (
	"os"
	"path/filepath"
)

// newWildcardEntry 本质还是compositeEntry无需新增数据类型
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //remove * 获取基础的目录
	var compositeEntry CompositeEntry

	//walkFn 遍历函数
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//如果文件是文件夹则跳过，通配符类路径不能递归匹配子目录下的jar
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if judgeIsJar(path) {
			zipEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, zipEntry)
		}
		return nil
	}
	//遍历baseDir
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
