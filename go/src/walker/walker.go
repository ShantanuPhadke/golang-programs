package walker

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	value := extractValue(x)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(value.MapIndex(key).Interface(), fn)
		}
	case reflect.String:
		fn(value.String())
	default:
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			switch field.Kind() {
			case reflect.String:
				fn(field.String())
			case reflect.Struct:
				walk(field.Interface(), fn)
			}
		}
	}

	/* for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
		// Otherwise if its a struct then go recursively into that struct
		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	} */
}

func extractValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	// Getting the actual object of a pointer
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}
