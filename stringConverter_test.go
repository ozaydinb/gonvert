package gonvert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type model struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestConvertString_ToIntCanConvertNumericValue(t *testing.T) {
	val, err := ConvertString("12345").ToInt()
	assert.NoError(t, err)
	assert.Equal(t, val, 12345)
}

func TestConvertString_ToIntCanConvertMinusNumericValue(t *testing.T) {
	val, err := ConvertString("-12345").ToInt()
	assert.NoError(t, err)
	assert.Equal(t, val, -12345)
}

func TestConvertString_ToIntReturnDefaultValue(t *testing.T) {
	val, err := ConvertString("123abc").ToInt(42)
	assert.NoError(t, err)
	assert.Equal(t, val, 42)
}

func TestConvertString_ToIntReturnError(t *testing.T) {
	val, err := ConvertString("123abc").ToInt()
	assert.Error(t, err)
	assert.Equal(t, val, 0)
}

func TestConvertString_ToInt32CanConvertNumericValue(t *testing.T) {
	val, err := ConvertString("12345").ToInt32()
	assert.NoError(t, err)
	assert.Equal(t, val, int32(12345))
}

func TestConvertString_ToInt32CanConvertMinusNumericValue(t *testing.T) {
	val, err := ConvertString("-12345").ToInt32()
	assert.NoError(t, err)
	assert.Equal(t, val, int32(-12345))
}

func TestConvertString_ToInt32ReturnDefaultValue(t *testing.T) {
	val, err := ConvertString("123abc").ToInt32(42)
	assert.NoError(t, err)
	assert.Equal(t, val, int32(42))
}

func TestConvertString_ToInt32ReturnError(t *testing.T) {
	val, err := ConvertString("123abc").ToInt32()
	assert.Error(t, err)
	assert.Equal(t, val, int32(0))
}

func TestConvertString_ToModelCanConvertValidJsonValue(t *testing.T) {
	m := &model{}
	success, err := ConvertString("{\"name\":\"Barış\", \"age\":1}").ToModel(m)
	assert.True(t, success)
	assert.NoError(t, err)
	assert.Equal(t, m.Name, "Barış")
	assert.Equal(t, m.Age, 1)
}

func TestConvertString_ToModelReturnError(t *testing.T) {
	m := &model{}
	success, err := ConvertString("it is not a json").ToModel(m)
	assert.False(t, success)
	assert.Error(t, err)
	assert.Equal(t, m.Name, "")
	assert.Equal(t, m.Age, 0)
}
