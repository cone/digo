package digo

import (
	"errors"
	"reflect"
)

type Injector struct{}

func (this *Injector) New(key string) (interface{}, error) {
	t, err := TypeRegistry.Get(key)
	if err != nil {
		return "", err
	}
	cp := reflect.New(t).Elem()
	return cp.Interface(), nil
}

func (this *Injector) Resolve(node *DependencyNode) (interface{}, error) {
	t, err := TypeRegistry.Get(node.Name)
	if err != nil {
		return struct{}{}, err
	}

	cp := reflect.New(t).Elem()

	for _, dependency := range node.Dependencies {
		f := cp.FieldByName(dependency.FieldName)

		if f.IsValid() {

			if f.CanSet() {

				depcp, err := this.Resolve(dependency)
				if err != nil {
					return struct{}{}, err
				}

				f.Set(reflect.ValueOf(depcp))

			} else {
				return struct{}{}, errors.New("Field cannot be set")
			}

		} else {
			return struct{}{}, errors.New("Invalid Field")
		}

	}

	return cp.Interface(), nil
}

// TODO: Add a 'Get' method that returns the
// struct with all its dependencies solved
