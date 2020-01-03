/*
@Time : 2020/1/3 17:16
@Software: GoLand
@File : init
@Author : Bingo <airplayx@gmail.com>
*/
package Map2struct

import (
	"errors"
	"fmt"
	"reflect"
)

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj ", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value ", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type ")
	}

	structFieldValue.Set(val)
	return nil
}

type MyStruct struct {
	Name string
	Age  int64
}

func (s *MyStruct) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
