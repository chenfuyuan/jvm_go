package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry //启动类路径（bootstrap classpath）
	extClasspath  Entry //外部类路径（extension classpath）
	userClasspath Entry //用户类路径（user classpath）
}

// parseBootAndExtClasspath 根据jre路径解析启动类路径和扩展类路径
func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	//jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)

	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

// parseUserClasspath 解析用户类路径
func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}

// 获取jre目录
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")
}

// exists 判断目录是否存在
func exists(path string) bool {
	//判断路径是否存在
	if _, err := os.Stat(path); err != nil {
		//处理错误
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Parse 解析classpath
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// ReadClass 根据className 依次从boot,ext,user类路径读取类
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, entry, err := c.bootClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := c.extClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}

	return c.userClasspath.ReadClass(className)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}
