package digo

import (
	//"fmt"
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

		depcp, err := this.Resolve(dependency)
		if err != nil {
			return struct{}{}, err
		}

		f.Set(reflect.ValueOf(depcp))
	}

	return cp.Interface(), nil
}

// TODO: Add a 'Get' method that returns the
// struct with all its dependencies solved
