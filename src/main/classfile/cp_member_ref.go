package classfile

type ConstantMemberRefInfo struct {
	pool             ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (info *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	info.classIndex = reader.readUint16()
	info.nameAndTypeIndex = reader.readUint16()
}

func (info *ConstantMemberRefInfo) ClassName() string {
	return info.pool.getClassName(info.classIndex)
}

func (info *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return info.pool.getNameAndType(info.nameAndTypeIndex)
}

// ConstantFieldRefInfo
/*
CONSTANT_FieldRef_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}*/
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

// ConstantMethodRefInfo
/*
CONSTANT_MethodRef_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

// ConstantInterfaceMethodRefInfo
/*
CONSTANT_InterfaceMethodRef_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}
