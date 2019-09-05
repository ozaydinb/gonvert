package gonvert

import (
	"encoding/json"
	"strconv"
)

//ConvertString used for converting string to any convertible type
func ConvertString(val string) StringConverter {
	return &stringConverterImp{
		val: val,
	}
}

//StringConverter used for converting string to any convertible type
type StringConverter interface {
	ToInt(defaultValue ...int) (int, error)
	ToInt32(defaultValue ...int32) (int32, error)
	ToStruct(valType interface{}) error
}

type stringConverterImp struct {
	val string
}

func (c *stringConverterImp) ToInt32(defaultValue ...int32) (int32, error) {
	result, err := strconv.ParseInt(c.val, 10, 32)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return 0, err
	}
	return int32(result), nil
}

func (c *stringConverterImp) ToInt(defaultValue ...int) (int, error) {
	result, err := strconv.Atoi(c.val)
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return result, err
}

func (c *stringConverterImp) ToStruct(valType interface{}) error {
	err := json.Unmarshal([]byte(c.val), valType)
	return err
}
