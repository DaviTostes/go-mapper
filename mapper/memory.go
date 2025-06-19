package mapper

import (
	"reflect"
	"sync"
)

var store sync.Map

func SaveStruct(value any) {
	key := getStructKey(value)
	store.Store(key, value)
}

func LoadStruct[T any]() (T, bool) {
	var zero T
	key := getStructKey(zero)

	v, ok := store.Load(key)
	if !ok {
		return zero, false
	}

	result, ok := v.(T)
	if !ok {
		return zero, false
	}

	return result, true
}

func getStructKey(value any) string {
	typ := reflect.TypeOf(value)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.PkgPath() + "." + typ.Name()
}
