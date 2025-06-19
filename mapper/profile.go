package mapper

import (
	"errors"
	"reflect"
)

type Profile[S, D any] struct {
	Source  S
	Destiny D
	Maps    map[string]func(src S) any
}

func (p *Profile[S, D]) ForMember(field string, fn func(src S) any) {
	p.Maps[field] = fn
}

func CreateProfile[S, D any](source S, destiny D) (Profile[S, D], error) {
	if reflect.TypeOf(source).Kind() != reflect.Struct {
		return Profile[S, D]{}, errors.New("Failed to create Profile: Source is not a struct")
	}

	if reflect.TypeOf(destiny).Kind() != reflect.Struct {
		return Profile[S, D]{}, errors.New("Failed to create Profile: Destiny is not a struct")
	}

	p := Profile[S, D]{Source: source, Destiny: destiny, Maps: make(map[string]func(src S) any)}
	SaveStruct(p)

	return p, nil
}
