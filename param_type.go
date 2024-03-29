package ssp

type ParamType int

const (
	ParamTypeByte ParamType = iota
	ParamTypeBool
	ParamTypeInt8
	ParamTypeUInt8
	ParamTypeInt16
	ParamTypeUInt16
	ParamTypeInt32
	ParamTypeUInt32
	ParamTypeInt64
	ParamTypeUInt64
	ParamTypeFloat
	//ParamTypeChar
	ParamTypeString
)
