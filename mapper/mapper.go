package mapper

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

func Map[S, D any](s S, d *D) error {
	valS := reflect.ValueOf(s)
	typS := valS.Type()

	valD := reflect.ValueOf(d).Elem()
	typD := valD.Type()

	p, ok := loadStruct[Profile[S, D]]()
	if !ok {
		return errors.New("Profile [" + typS.Name() + "] -> [" + typD.Name() + "] not found")
	}

	var wg sync.WaitGroup
	errs := make(chan error, 2)

	wg.Add(2)

	go func() {
		for i := range typS.NumField() {
			sField := typS.Field(i)
			if sField.Type.Kind() == reflect.Struct {
				continue
			}

			dField := valD.FieldByName(sField.Name)
			vs := valS.Field(i)

			if err := checkAssignability(dField, vs.Interface()); err != nil {
				errs <- err
			}

			dField.Set(vs)
		}
		wg.Done()
	}()

	go func() {
		for field, fn := range p.Maps {
			value := fn(s)
			dField := valD.FieldByName(field)

			if err := checkAssignability(dField, value); err != nil {
				errs <- err
			}

			dField.Set(reflect.ValueOf(value))
		}
		wg.Done()
	}()

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func checkAssignability(dField reflect.Value, dValue any) error {
	if !dField.IsValid() {
		return errors.New("Field " + dField.String() + " has no value")
	}
	if !dField.CanSet() {
		return errors.New("Field " + dField.String() + " can't be setted")
	}
	if !reflect.TypeOf(dValue).AssignableTo(dField.Type()) {
		return errors.New(
			fmt.Sprint("Cannot assign value '", dValue, "' to field ", dField.String()),
		)
	}

	return nil
}
