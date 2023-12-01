package validator

import "testing"

type Person struct {
	Name string `validate:"required" label:"姓名"`
	Age  int
}

func TestStruct(t *testing.T) {
	p := Person{}
	err := Struct(p)
	if err != nil {
		t.Log(err)
	}
}
