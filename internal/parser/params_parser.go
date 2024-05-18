package parser

import (
	"errors"
	"yesbotics/ssp/internal/types"
	"yesbotics/ssp/pkg/config"
)

type ParamsParser struct {
	params       []types.Type
	currentParam types.Type
	paramIndex   int
	paramCount   int
	hasParams    bool
}

func NewParamsParser(paramTypes []config.ParamType) *ParamsParser {

	//var test = params.typeByte{}
	//var typeInstances []*params.Type = make([]*params.Type, 0)
	//var typeInstances []*params.Type
	var params []types.Type

	for _, paramType := range paramTypes {
		switch paramType {
		case config.ParamTypeByte:
			params = append(params, types.NewTypeByte())
		case config.ParamTypeBool:
			params = append(params, types.NewTypeBool())
		case config.ParamTypeInt8:
			params = append(params, types.NewTypeInt8())
		case config.ParamTypeUint8:
			params = append(params, types.NewTypeUint8())
		case config.ParamTypeInt16:
			params = append(params, types.NewTypeInt16())
		case config.ParamTypeUint16:
			params = append(params, types.NewTypeUint16())
		case config.ParamTypeInt32:
			params = append(params, types.NewTypeInt32())
		case config.ParamTypeUint32:
			params = append(params, types.NewTypeUint32())
		case config.ParamTypeInt64:
			params = append(params, types.NewTypeInt64())
		case config.ParamTypeUint64:
			params = append(params, types.NewTypeUint64())
		case config.ParamTypeFloat32:
			params = append(params, types.NewTypeFloat32())
		case config.ParamTypeFloat64:
			params = append(params, types.NewTypeFloat64())
		case config.ParamTypeChar:
			params = append(params, types.NewTypeChar())
		case config.ParamTypeString:
			params = append(params, types.NewTypeString())
		}
	}

	var hasParams = len(params) > 0
	var firstParam types.Type = nil

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
	if !r.hasParams {
		return errors.New("tried to add byte to params but no params defined")
	}

	if r.currentParam.IsFull() {
		r.paramIndex++
		if r.paramIndex >= r.paramCount {
			return errors.New("tried to add byte to param parser but all types filled")
		}

		r.currentParam = r.params[r.paramIndex]
	}

	r.currentParam.AddByte(bite)

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
			return r.currentParam.IsFull()
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
			//param.(types.Type[uint8]).Reset()
			param.Reset()
		}
	}
}

func (r *ParamsParser) GetData() ([]any, error) {
	var data []any
	if r.hasParams {
		for _, param := range r.params {
			//data = append(r.getDataSwitched(param))
			paramData, err := param.GetData()
			if err != nil {
				return nil, err
			}
			data = append(data, paramData)
		}
	}

	return data, nil
}

//func (r *ParamsParser) getDataSwitched(param any) {
//    switch typeClass := param.(type) {
//    case types.typeByte:
//        typeClass.
//    }
//}

func (r *ParamsParser) Dispose() {

}
