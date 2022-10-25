package classfile

type MemberInfo struct {
	constantPool    ConstantPool //常量池
	accessFlags     uint16       //访问权限
	nameIndex       uint16       //
	descriptorIndex uint16       //
	attributes      []AttributeInfo
}

// readMembers 读取所有成员信息
func readMembers(reader *ClassReader, constantPool ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	result := make([]*MemberInfo, memberCount)
	for i := range result {
		result[i] = readMember(reader, constantPool)
	}
	return result
}

// readMember 读取单个成员信息
func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    constantPool,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		//attributes: readAttributes(reader,constantPool),
	}
}
