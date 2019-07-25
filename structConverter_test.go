package gonvert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
	Name string
	Age  int
}

func TestStructConverter_ToJsonReturnString(t *testing.T) {
	//Given
	model := &testStruct{
		Name: "Barış",
		Age:  1,
	}

	//When
	jsonString, err := ConvertStruct(model).ToJson()

	//Then
	assert.NoError(t, err)
	assert.Equal(t, "{\"Name\":\"Barış\",\"Age\":1}", jsonString)
}

func TestStructConverter_ToJsonReturnError(t *testing.T) {
	c := make(chan bool)
	_, err := ConvertStruct(c).ToJson()
	assert.Error(t, err)
}

func TestStructConverter_ToMapReturnMap(t *testing.T) {
	model := &testStruct{
		Name: "Barış",
		Age:  1,
	}

	m, err := ConvertStruct(model).ToMap()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(m))
}

func TestStructConverter_ToMapReturnError(t *testing.T) {
	c := make(chan bool)
	_, err := ConvertStruct(c).ToMap()

	assert.Error(t, err)
}
