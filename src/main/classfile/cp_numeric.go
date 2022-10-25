package classfile

import "math"

// ConstantIntegerInfo 存放boolean、byte、short、char、int32
/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (info *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	info.val = int32(bytes)
}

func (info *ConstantIntegerInfo) Value() int32 {
	return info.val
}

// ConstantFloatInfo
/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (info *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	info.val = math.Float32frombits(bytes)
}

func (info *ConstantFloatInfo) Value() float32 {
	return info.val
}

// ConstantLongInfo
/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val int64
}

func (info *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	info.val = int64(bytes)
}

func (info *ConstantLongInfo) Value() int64 {
	return info.val
}

// ConstantDoubleInfo
/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (info *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	info.val = math.Float64frombits(bytes)
}

func (info *ConstantDoubleInfo) Value() float64 {
	return info.val
}
