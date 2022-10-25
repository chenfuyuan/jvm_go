package classfile

import "fmt"

// ConstantPool 常量池
type ConstantPool struct {
	count int
	data  []ConstantInfo
}

// readConstantPool 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	//读取常量池长度
	cpCount := int(reader.readUint16())
	data := make([]ConstantInfo, cpCount)
	pool := ConstantPool{count: cpCount, data: data}
	for i := 1; i < cpCount; i++ {
		data[i] = readConstantInfo(reader, pool)
		switch data[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return pool
}

func (pool ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := pool.data[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("invalid constant pool index: %v", index))
}

func (pool ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := pool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := pool.getUtf8(ntInfo.nameIndex)
	_type := pool.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (pool ConstantPool) getClassName(index uint16) string {
	classInfo := pool.getConstantInfo(index).(*ConstantClassInfo)
	return pool.getUtf8(classInfo.nameIndex)
}

func (pool ConstantPool) getUtf8(index uint16) string {
	utf8Info := pool.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
