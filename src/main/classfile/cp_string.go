package classfile

// ConstantStringInfo
/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	pool        ConstantPool
	stringIndex uint16 //常量池索引
}

func (info *ConstantStringInfo) readInfo(reader *ClassReader) {
	info.stringIndex = reader.readUint16()
}

func (info *ConstantStringInfo) String() string {
	return info.pool.getUtf8(info.stringIndex)
}
