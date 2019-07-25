package gonvert

import (
	"encoding/json"
	"strconv"
)

func ConvertString(val string) stringConverter {
	return &stringConverterImp{
		val: val,
	}
}

type stringConverter interface {
	ToInt32(defaultValue ...int32) (int32, error)
	ToInt(defaultValue ...int) (int, error)
	ToModel(valType interface{}) (bool, error)
}
type stringConverterImp struct {
	val string
}

func (c *stringConverterImp) ToInt32(defaultValue ...int32) (int32, error) {
	result, err := strconv.ParseInt(c.val, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

func (c *stringConverterImp) ToInt(defaultValue ...int) (int, error) {
	result, err := strconv.Atoi(c.val)
	return result, err
}

func (c *stringConverterImp) ToModel(valType interface{}) (bool, error) {
	err := json.Unmarshal([]byte(c.val), valType)
	if err != nil {
		return false, err
	}
	return true, nil
}
