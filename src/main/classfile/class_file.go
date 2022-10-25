package classfile

import "fmt"

const JAVA_CLASSFILE_MAGIC = 0xCAFEBABE //class文件标识

type ClassFile struct {
	magic           uint32 // 魔数 0xCAFEBABE   class文件标识
	minorVersion    uint16
	majorVersion    uint16
	constantPool    ConstantPool //常量池
	accessFlags     uint16
	thisClass       uint16
	superClass      uint16
	interfacesCount uint16
	interfaces      []uint16
	fieldsCount     uint16
	//fields []*MemberInfo
	methodsCount uint16
	//methods []*MemberInfo
	attributesCount uint16
	//attributes  []AttributeInfo
}

// Parse 解析classData
// @param classData 类数据
// @return classFile 类文件
// @return err 错误
func Parse(classData []byte) (classFile *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf := &ClassFile{}
	cf.read(cr)
	return cf, nil
}

// read 读取数据，并解析
// @param reader 类读取工具
func (cf *ClassFile) read(reader *ClassReader) {
	//依次读取class文件，进行解析
	cf.readAndCheckMagic(reader)               //读取和校验魔数
	cf.readAndCheckVersion(reader)             //读取和校验版本号
	cf.constantPool = readConstantPool(reader) //常量池解析
	cf.readAccessFlags(reader)                 //访问权限解析
	cf.readThisClass(reader)                   //本类常量池索引
	cf.readSuperClass(reader)                  //超类常量池索引
	cf.readInterfaces(reader)                  //接口索引表
	//cf.readFields(reader)//字段
	//cf.readMethods(reader)//方法表
	//cf.readAttributes(reader)//属性值
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != JAVA_CLASSFILE_MAGIC {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

// readAndCheckVersion 校验class的版本号，判断本jvm是否支持该版本
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// readAccessFlags 读取访问权限
func (cf *ClassFile) readAccessFlags(reader *ClassReader) {
	cf.accessFlags = reader.readUint16()
}

// readThisClass 读取本类在常量池中的索引
func (cf *ClassFile) readThisClass(reader *ClassReader) {
	cf.thisClass = reader.readUint16()
}

// readSuperClass 读取超类在常量池中的索引
func (cf *ClassFile) readSuperClass(reader *ClassReader) {
	cf.superClass = reader.readUint16()
}

// readInterfaces 读取接口索引
func (cf *ClassFile) readInterfaces(reader *ClassReader) {
	cf.interfaces = reader.readUint16s()
	cf.interfacesCount = uint16(len(cf.interfaces))
}

/*================Getter===============*/

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

//todo class_file field
/*
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}
*/
