package digo

import (
	"errors"
	"reflect"
)

type ContextManager struct {
	contexts map[string]*Context
}

func (this *ContextManager) Context(filePath string) (*Context, error) {
	if ctx, exists := this.contexts[filePath]; exists {
		return ctx, nil
	}

	ctx, err := this.newContext(filePath)
	if err != nil {
		return nil, err
	}

	this.contexts[filePath] = ctx

	return ctx, nil

}

func (this *ContextManager) newContext(filePath string) (*Context, error) {
	ctx := &Context{}

	err := ctx.unmarshal(filePath)
	if err != nil {
		return nil, errors.New("Error creating new Context -> " + err.Error())
	}

	return ctx, nil
}

func (this *ContextManager) New(key string, isPtr bool) (interface{}, error) {
	t, err := TypeRegistry.Get(key)
	if err != nil {
		return struct{}{}, errors.New("Error getting the type from TypeRegistry -> " + err.Error())
	}

	if isPtr {
		return reflect.New(t).Interface(), nil
	}

	return reflect.New(t).Elem().Interface(), nil
}
