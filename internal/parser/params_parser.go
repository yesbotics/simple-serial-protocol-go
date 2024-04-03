package parser

import (
	"errors"
	"yesbotics/ssp/internal/types"
	"yesbotics/ssp/pkg/config"
)

type ParamsParser struct {
	params       []any
	currentParam any
	paramIndex   int
	paramCount   int
	hasParams    bool
}

func NewParamsParser(paramTypes []config.ParamType) *ParamsParser {

	//var test = params.TypeByte{}
	//var typeInstances []*params.Type = make([]*params.Type, 0)
	//var typeInstances []*params.Type
	var params []any

	for _, paramType := range paramTypes {
		switch paramType {
		case config.ParamTypeByte:
			params = append(params, &(types.TypeByte{}))
		case config.ParamTypeBool:
			params = append(params, &(types.TypeBool{}))
		case config.ParamTypeInt8:
			params = append(params, &(types.TypeInt8{}))
		case config.ParamTypeUint8:
			params = append(params, &(types.TypeUint8{}))
		case config.ParamTypeInt16:
			params = append(params, &(types.TypeInt16{}))
		case config.ParamTypeUint16:
			params = append(params, &(types.TypeUint16{}))
		case config.ParamTypeInt32:
			params = append(params, &(types.TypeInt32{}))
		case config.ParamTypeUint32:
			params = append(params, &(types.TypeUint32{}))
		case config.ParamTypeInt64:
			params = append(params, &(types.TypeInt64{}))
		case config.ParamTypeUint64:
			params = append(params, &(types.TypeUint64{}))
		case config.ParamTypeFloat32:
			params = append(params, &(types.TypeFloat32{}))
		case config.ParamTypeFloat64:
			params = append(params, &(types.TypeFloat64{}))
		case config.ParamTypeChar:
			params = append(params, &(types.TypeChar{}))
		case config.ParamTypeString:
			params = append(params, &(types.TypeString{}))
		}
	}

	var hasParams = len(params) > 0
	var firstParam any = nil

	if hasParams {
		firstParam = params[0]
	}

	return &ParamsParser{
		params:       params,
		currentParam: firstParam,
		hasParams:    hasParams,
		paramIndex:   0,
		paramCount:   len(params),
	}
}

func (r *ParamsParser) AddByte(bite byte) error {
	if r.hasParams {
		return errors.New("tried to add byte to params but no params defined")
	}

	param := r.currentParam.(types.Type[uint8])

	if param.IsFull() {
		r.paramIndex++
		if r.paramIndex >= r.paramCount {
			return errors.New("tried to add byte to param parser but all types filled")
		}

		r.currentParam = r.params[r.paramIndex]
	}

	param.AddByte(bite)

	return nil
}

func (r *ParamsParser) IsFull() bool {
	if r.hasParams {
		if r.paramIndex < (r.paramCount - 1) {
			/**
			 * Not reached last type
			 */
			return false
		} else {
			/**
			 * Last param filled?
			 */
			return r.currentParam.(types.Type[uint8]).IsFull()
		}
	} else {
		/**
		 * No types defined -> always full
		 */
		return true
	}
}

func (r *ParamsParser) Reset() {
	if r.hasParams {
		for _, param := range r.params {
			param.(types.Type[uint8]).Reset()
		}
	}
}

func (r *ParamsParser) GetData() []any {
	var data []any
	if r.hasParams {
		for _, param := range r.params {
			//data = append(r.getDataSwitched(param))
			data = append(data, param)
		}
	}

	return data
}

//func (r *ParamsParser) getDataSwitched(param any) {
//    switch typeClass := param.(type) {
//    case types.TypeByte:
//        typeClass.
//    }
//}

func (r *ParamsParser) Dispose() {

}
