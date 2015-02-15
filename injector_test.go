package digo

import (
	"reflect"
	"testing"
)

func TestInjector_New(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Type assertion failed!")
		}
	}()

	injector := new(Injector)

	test := "hello"

	TypeRegistry["string"] = reflect.TypeOf(test)

	cp, err := injector.New("string")
	if err != nil {
		t.Error("Type not found")
	}

	test = cp.(string)
}
