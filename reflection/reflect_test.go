package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFieldValueByType(t *testing.T) {
	type testStruct2 struct {
		A string
		B int
	}

	type testStruct struct {
		Name string
		Age  int
		Test testStruct2
	}

	ts := testStruct{
		Name: "test",
		Age:  1,
		Test: testStruct2{},
	}

	fields := GetFieldValueByType(ts, reflect.TypeOf(testStruct2{}))	
	fmt.Println(fields)	
}