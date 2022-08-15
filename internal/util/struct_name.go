package util

import "reflect"

// StructName receives a struct and return the struct's name.
func StructName(s any) string {
	if t := reflect.TypeOf(s); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
