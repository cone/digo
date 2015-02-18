package digo

import (
	"fmt"
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

	if t.Kind() == reflect.Struct {
		for _, dep := range node.Dependencies {
			resolvedDep, err := this.Resolve(dep)
			if err != nil {
				return struct{}{}, err
			}
			fmt.Printf("The node %s has a dep of type %v", node.Name, reflect.TypeOf(resolvedDep))
			fmt.Println()
		}
	}
	//call resolve for each of its dependencies
	//iterate the struct fields (if it is a struct)
	//and assign the values to the correct field

	return cp.Interface(), nil
}
