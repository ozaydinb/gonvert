package gonvert

import "encoding/json"

type StructConverter interface {
	ToJson() (string, error)
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
	result, err := json.Marshal(c.val)
	if err != nil {
		return "", err
	}
	jsonResult := string(result)
	return jsonResult, nil
}
