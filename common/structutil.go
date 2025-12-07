package common

import "reflect"

func CloneEmpty(obj interface{}) interface{} {
	return reflect.New(reflect.TypeOf(obj).Elem()).Interface()
}

func SetErrorAndErrorMessage(response interface{}, errCode *int32, errMessage *string) {
	SetFieldValueToStruct(response, "Error", errCode)
	SetFieldValueToStruct(response, "ErrorMessage", errMessage)
}

func SetFieldValueToStruct(obj interface{}, name string, val interface{}) {
	objValue := reflect.ValueOf(obj)
	for objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	field := objValue.FieldByName(name)
	if !field.IsValid() {
		return
	}

	if !field.CanSet() {
		return
	}

	field.Set(reflect.ValueOf(val))
}
