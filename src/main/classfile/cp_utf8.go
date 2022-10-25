package classfile

import "go_jvm/src/main/utils"

type ConstantUtf8Info struct {
	str string
}

func (info *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	info.str = utils.ByteToString(bytes)
}

func (info *ConstantUtf8Info) Str() string {
	return info.str
}
