package utils

import (
	"errors"
	"reflect"
)

func CopySliceVoidly(src interface{}) (interface{}, error) {
	if reflect.TypeOf(src).Kind() != reflect.Slice {
		return nil, errors.New("src must be a kind of slice")
	}
	elemType := reflect.TypeOf(src)
	return reflect.MakeSlice(elemType, 0, 0).Interface(), nil
}

func CopyInstanceVoidly(src interface{}) (interface{}, error) {
	if reflect.TypeOf(src).Kind() == reflect.Slice {
		return CopySliceVoidly(src)
	} else if reflect.TypeOf(src).Kind() == reflect.Interface {
		elemType := reflect.TypeOf(src)
		return reflect.New(elemType).Elem(), nil
	} else {
		return nil, errors.New("src must be a kind of slice")
	}
}
