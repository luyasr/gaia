package reflection

import (
	"testing"
)

func TestStructToMap(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int
	}

	obj := testStruct{
		Name: "Luya",
		// Age:  20,
	}

	m := StructToMap(obj)
	rm := StructToMapRmEmpty(obj)

	t.Log(m, rm)
}