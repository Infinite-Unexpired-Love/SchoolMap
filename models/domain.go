package models

import "reflect"

type Updatable interface {
	Update(updatable interface{})
}

func updateStructFields(target interface{}, item interface{}) {
	updatableValue := reflect.ValueOf(item)
	if updatableValue.Kind() == reflect.Ptr {
		updatableValue = updatableValue.Elem()
	}

	targetValue := reflect.ValueOf(target).Elem()

	for i := 0; i < updatableValue.NumField(); i++ {
		field := updatableValue.Field(i)
		fieldType := updatableValue.Type().Field(i)
		fieldName := fieldType.Name

		targetField := targetValue.FieldByName(fieldName)
		if !targetField.IsValid() || !targetField.CanSet() {
			continue
		}

		if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			targetField.Set(field)
		}
	}
}

type Cascade interface {
	SetParentID(*uint)
	GetParentID() *uint
	GetID() *uint
}
