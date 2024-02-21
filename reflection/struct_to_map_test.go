package reflection

import (
	"testing"
)

func TestStructToMap(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  *int
		C    struct {
			Addr string `json:"addr"`
		}
		D int `json:"d"`
	}

	// age := 20

	obj := testStruct{
		// Name: "Luya",
		// Age:  &age,
		// C: struct {
		// 	Addr string `json:"addr"`
		// }{Addr: "127.0.0.1:80"},
	}

	m := StructToMap(obj)

	t.Log(m)
}
