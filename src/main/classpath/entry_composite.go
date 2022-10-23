package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var result []Entry
	//通过分隔符对pathList进行分割，分隔出path
	for _, path := range strings.Split(pathList, PathListSeparator) {
		//根据path，构建entry，并加入集合
		entry := newEntry(path)
		result = append(result, entry)
	}
	return result
}

func (c CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, from, err := entry.ReadClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (c CompositeEntry) String() string {
	strs := make([]string, len(c))
	for i, entry := range c {
		strs[i] = entry.String()
	}
	return strings.Join(strs, PathListSeparator)
}
