package digo

import (
	"errors"
	"fmt"
	"reflect"
)

type TypeMap map[string]reflect.Type

func (this TypeMap) Add(param interface{}) error {
	t := reflect.TypeOf(param)
	key := fmt.Sprintf("%v", t)
	this[key] = t
	return nil
}

func (this TypeMap) AddType(param reflect.Type) error {
	key := fmt.Sprintf("%v", param)
	this[key] = param
	return nil
}

func (this TypeMap) Get(key string) (reflect.Type, error) {
	if t, exists := this[key]; !exists {
		return t, errors.New("No such Type: " + key)
	}
	return this[key], nil
}
