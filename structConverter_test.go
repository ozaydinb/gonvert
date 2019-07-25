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
