package digo

import (
	"errors"
	"reflect"
)

func init() {
	if depInjector == nil {
		depInjector = new(Injector)
	}
}

var depInjector *Injector

type Injector struct{}

func (this *Injector) New(key string) (interface{}, error) {
	cp, err := this.newTypeOf(key)
	if err != nil {
		return struct{}{}, errors.New("Error creating new Type -> " + err.Error())
	}

	return cp.Interface(), nil
}

func (this *Injector) Resolve(node *NodeData, nodeMap map[string]*NodeData) (interface{}, error) {
	cp, err := this.newTypeOf(node.Type)
	if err != nil {
		return struct{}{}, errors.New("Error creating new Type -> " + err.Error())
	}

	for _, dependency := range node.Deps {

		if depNode, exists := nodeMap[dependency.ID]; exists {

			err := this.assignValues(cp, depNode, nodeMap)
			if err != nil {
				return struct{}{}, errors.New("Error resolving dependencies -> " + err.Error())
			}

		} else {
			return struct{}{}, errors.New("Dependency Id not found -> " + err.Error())
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

func (this *Injector) assignValues(cp reflect.Value, dependency *NodeData, nodeMap map[string]*NodeData) error {
	f := cp.FieldByName(dependency.Field)

	if f.IsValid() {

		if f.CanSet() {

			depcp, err := this.Resolve(dependency, nodeMap)
			if err != nil {
				return err
			}

			f.Set(reflect.ValueOf(depcp))

		} else {
			return errors.New("Field cannot be set")
		}

	} else {
		return errors.New("Invalid Field: " + dependency.Field)
	}

	return nil
}
