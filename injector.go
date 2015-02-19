package digo

import (
	"errors"
	"reflect"
)

type Injector struct{}

func (this *Injector) New(key string) (interface{}, error) {
	cp, err := this.newTypeOf(key)
	if err != nil {
		return struct{}{}, errors.New("Error creating new Type -> " + err.Error())
	}

	return cp.Interface(), nil
}

func (this *Injector) Resolve(node *DependencyNode) (interface{}, error) {
	cp, err := this.newTypeOf(node.Name)
	if err != nil {
		return struct{}{}, errors.New("Error creating new Type -> " + err.Error())
	}

	for _, dependency := range node.Dependencies {

		err := this.assignValues(cp, dependency)
		if err != nil {
			return struct{}{}, errors.New("Error resolving dependencies -> " + err.Error())
		}

	}

	return cp.Interface(), nil
}

func (this *Injector) newTypeOf(key string) (reflect.Value, error) {
	t, err := TypeRegistry.Get(key)
	if err != nil {
		return reflect.Value{}, errors.New("Error getting the type from TypeRegistry -> " + err.Error())
	}

	return reflect.New(t).Elem(), nil
}

func (this *Injector) assignValues(cp reflect.Value, dependency *DependencyNode) error {
	f := cp.FieldByName(dependency.FieldName)

	if f.IsValid() {

		if f.CanSet() {

			depcp, err := this.Resolve(dependency)
			if err != nil {
				return err
			}

			f.Set(reflect.ValueOf(depcp))

		} else {
			return errors.New("Field cannot be set")
		}

	} else {
		return errors.New("Invalid Field")
	}

	return nil
}

// TODO: Add a 'Get' method that returns the
// struct with all its dependencies solved
