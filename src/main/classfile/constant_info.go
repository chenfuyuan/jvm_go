package classfile

const (
	CONSTANT_UTF8                 = 1
	CONSTANT_INTEGER              = 3
	CONSTANT_FLOAT                = 4
	CONSTANT_LONG                 = 5
	CONSTANT_DOUBLE               = 6
	CONSTANT_CLASS                = 7
	CONSTANT_STRING               = 8
	CONSTANT_FIELD_REF            = 9
	CONSTANT_METHOD_REF           = 10
	CONSTANT_INTERFACE_METHOD_REF = 11
	CONSTANT_NAME_AND_TYPE        = 12

	CONSTANT_METHOD_HANDLE = 15
	CONSTANT_METHOD_TYPE   = 16

	CONSTANT_INVOKE_DYNAMIC = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// newConstantInfo 根据tag新建变量
func newConstantInfo(tag uint8, pool ConstantPool) ConstantInfo {
	//根据tag判断返回什么类型的变量
	switch tag {
	case CONSTANT_UTF8:
		return &ConstantUtf8Info{}
	case CONSTANT_INTEGER:
		return &ConstantIntegerInfo{}
	case CONSTANT_FLOAT:
		return &ConstantFloatInfo{}
	case CONSTANT_LONG:
		return &ConstantLongInfo{}
	case CONSTANT_DOUBLE:
		return &ConstantDoubleInfo{}
	case CONSTANT_CLASS:
		return &ConstantClassInfo{pool: pool}
	case CONSTANT_STRING:
		return &ConstantStringInfo{pool: pool}
	case CONSTANT_FIELD_REF:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{pool: pool}}
	case CONSTANT_METHOD_REF:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{pool: pool}}
	case CONSTANT_INTERFACE_METHOD_REF:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{pool: pool}}
	case CONSTANT_NAME_AND_TYPE:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_METHOD_HANDLE:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_METHOD_TYPE:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_INVOKE_DYNAMIC:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

// readConstantInfo 读取ConstantInfo
func readConstantInfo(reader *ClassReader, pool ConstantPool) ConstantInfo {
	//第一个字节是类型标识
	tag := reader.readUint8()
	info := newConstantInfo(tag, pool)
	//读取常量
	info.readInfo(reader)
	return info
}
