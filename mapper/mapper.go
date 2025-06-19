package mapper

import (
	"errors"
	"reflect"
)

func Map[S, D any](s S, d *D) error {
	valS := reflect.ValueOf(s)
	if valS.Kind() != reflect.Struct {
		return errors.New("Source is not a struct")
	}
	typS := valS.Type()

	typDP := reflect.TypeOf(d)
	if typDP.Kind() != reflect.Ptr {
		return errors.New("Destiny is not a pointer to a struct")
	}

	valD := reflect.ValueOf(d).Elem()
	if valD.Kind() != reflect.Struct {
		return errors.New("Destiny pointer is not to a struct")
	}
	typD := valD.Type()

	p, ok := LoadStruct[Profile[S, D]]()
	if !ok {
		return errors.New("Profile [" + typS.Name() + "] -> [" + typD.Name() + "] not found")
	}

	for field, fn := range p.Maps {
		value := fn(s)
		dField := valD.FieldByName(field)
		if dField.IsValid() && dField.CanSet() &&
			reflect.TypeOf(value).AssignableTo(dField.Type()) {
			dField.Set(reflect.ValueOf(value))
		}
	}

	for i := range typS.NumField() {
		sField := typS.Field(i)
		name := sField.Name

		dField := valD.FieldByName(name)
		if dField.IsValid() && dField.CanSet() && dField.Type() == sField.Type {
			vs := valS.Field(i)

			if !reflect.DeepEqual(vs.Interface(), dField.Interface()) {
				dField.Set(vs)
			}
		}
	}

	return nil
}
