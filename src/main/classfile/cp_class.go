package classfile

// ConstantClassInfo 类，超类索引 以及接口表中的接口索引指向的都是CONSTANT_CLASS_INFO常量
type ConstantClassInfo struct {
	pool      ConstantPool //常量池
	nameIndex uint16       //常量池中的索引
}

func (info *ConstantClassInfo) readInfo(reader *ClassReader) {
	info.nameIndex = reader.readUint16()
}
