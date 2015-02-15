package digo

import "reflect"

type Injector struct{}

func (this *Injector) New(key string) (interface{}, error) {
	t, err := TypeRegistry.Get(key)
	if err != nil {
		return "", err
	}
	cp := reflect.New(t).Elem()
	return cp.Interface(), nil
}
