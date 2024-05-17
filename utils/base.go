package utils

import (
	"errors"
	"reflect"
)

func GetVoidSlice(src interface{}) (interface{}, error) {
	if reflect.TypeOf(src).Kind() != reflect.Slice {
		return nil, errors.New("src must be a kind of slice")
	}
	elemType := reflect.TypeOf(src)
	return reflect.MakeSlice(elemType, 0, 0).Interface(), nil
}

func GetVoidInstance(src interface{}) (interface{}, error) {
	return nil, nil
}
