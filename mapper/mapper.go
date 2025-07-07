package mapper

import (
	"errors"
	"fmt"
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

	p, ok := loadStruct[Profile[S, D]]()
	if !ok {
		return errors.New("Profile [" + typS.Name() + "] -> [" + typD.Name() + "] not found")
	}

	for i := range typS.NumField() {
		sField := typS.Field(i)
		name := sField.Name
		dField := valD.FieldByName(name)

		if !dField.IsValid() {
			return errors.New("Field " + dField.String() + " has no value")
		}
		if !dField.CanSet() {
			return errors.New("Field " + dField.String() + " can't be setted")
		}

		if sField.Type.Kind() == reflect.Struct {
			continue
		}

		if dField.Type() != sField.Type {
			return errors.New(
				"Field " + dField.String() + " has different type of field " + sField.Name,
			)
		}

		vs := valS.Field(i)
		if !reflect.DeepEqual(vs.Interface(), dField.Interface()) {
			dField.Set(vs)
		}
	}

	for field, fn := range p.Maps {
		value := fn(s)
		dField := valD.FieldByName(field)

		if !dField.IsValid() {
			return errors.New("Field " + dField.String() + " has no value")
		}
		if !dField.CanSet() {
			return errors.New("Field " + dField.String() + " can't be setted")
		}
		if !reflect.TypeOf(value).AssignableTo(dField.Type()) {
			return errors.New(
				fmt.Sprint("Cannot assign value '", value, "' to field ", dField.String()),
			)
		}

		dField.Set(reflect.ValueOf(value))
	}

	return nil
}
