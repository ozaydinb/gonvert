package gonvert

import (
	"encoding/json"
	"errors"
	"github.com/fatih/structs"
)

type StructConverter interface {
	ToJson() (string, error)
	ToMap() (map[string]interface{}, error)
}

type StructConverterImp struct {
	val interface{}
}

func ConvertStruct(val interface{}) StructConverter {
	return &StructConverterImp{
		val: val,
	}
}

func (c *StructConverterImp) ToJson() (string, error) {
	isStruct := structs.IsStruct(c.val)
	if !isStruct {
		return "", errors.New("only structs can be converted")
	}
	result, err := json.Marshal(c.val)
	if err != nil {
		return "", err
	}
	jsonResult := string(result)
	return jsonResult, nil
}

func (c *StructConverterImp) ToMap() (map[string]interface{}, error) {
	isStruct := structs.IsStruct(c.val)
	if !isStruct {
		return nil, errors.New("only structs can be converted")
	}
	m := structs.Map(c.val)
	return m, nil
}
