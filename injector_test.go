package digo

import (
	"reflect"
	"testing"
)

func TestInjector_New(t *testing.T) {
	injector := new(Injector)

	test := "hello"

	TypeRegistry["string"] = reflect.TypeOf(test)

	cp, err := injector.New("string")
	if err != nil {
		t.Error("Type not found")
	}

	if _, ok := cp.(string); !ok {
		t.Error("Type assertion failed!")
	}
}

func TestInjector_Resolve(t *testing.T) {
	injector := new(Injector)

	test := "hello"

	TypeRegistry["digo.Dummy"] = reflect.TypeOf(Dummy{})
	TypeRegistry["string"] = reflect.TypeOf(test)

	dependencyTree := &DependencyNode{
		Name: "digo.Dummy",
		Dependencies: []*DependencyNode{
			{Name: "string"},
		},
	}

	target, err := injector.Resolve(dependencyTree)
	if err != nil {
		t.Error("The error has ocurred", err)
	}

	if _, ok := target.(Dummy); !ok {
		t.Error("Type assertion failed!")
	}

	t.Error()
}
