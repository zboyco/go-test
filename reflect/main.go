package main

import (
	"errors"
	"fmt"
	"reflect"
)

func appendToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()

	value.Set(reflect.Append(value, reflect.ValueOf(3)))

	fmt.Println(value)
	fmt.Println(value.Len())
}

func main() {
	arr := []int{1, 2}

	appendToSlice(&arr)

	fmt.Println(arr)

	m := &TestTable{
		Name: "TestName111222",
		Age:  18,
	}
	fmt.Println(validateStructType(m, "TestTable"))
	getStructColumnValue(&m, "Age")
}

type TestTable struct {
	Name string
	Age  int64
}

func validateStructType(obj interface{}, structName string) error {
	objType := reflect.TypeOf(obj)
	for objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	if objType.Kind() != reflect.Struct {
		return errors.New("param wrong type")
	}
	if objType.Name() != structName {
		return errors.New("param type is not main struct")
	}
	return nil
}

func getStructColumnValue(obj interface{}, columnName string) reflect.Value {
	valueType := reflect.ValueOf(obj)
	for valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}
	return valueType.FieldByName(columnName)
}
