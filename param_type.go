package ssp

type ParamType int

const (
	ParamTypeByte ParamType = iota
	ParamTypeBool
	ParamTypeInt8
	ParamTypeUint8
	ParamTypeInt16
	ParamTypeUint16
	ParamTypeInt32
	ParamTypeUint32
	ParamTypeInt64
	ParamTypeUint64
	ParamTypeFloat32
	ParamTypeFloat64
	ParamTypeChar
	ParamTypeString
)
