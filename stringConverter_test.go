package gonvert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type model struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestConvertString(t *testing.T) {
	val, err := ConvertString("12345").ToInt()
	assert.NoError(t, err)
	assert.Equal(t, val, 12345)
}

func TestStringConverterImp_ToModel(t *testing.T) {
	m := &model{}
	success, err := ConvertString("{\"name\":\"Barış\", \"age\":1}").ToModel(m)
	assert.True(t, success)
	assert.NoError(t, err)
	assert.Equal(t, m.Name, "Barış")
	assert.Equal(t, m.Age, 1)
}
