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
	TypeRegistry["digo.Dummy2"] = reflect.TypeOf(Dummy2{})

	dependencyTree := &DependencyNode{
		Name: "digo.Dummy",
		Dependencies: []*DependencyNode{
			{
				Name:      "digo.Dummy2",
				FieldName: "Field2",
			},
		},
	}

	target, err := injector.Resolve(dependencyTree)
	if err != nil {
		t.Error("The error has ocurred", err)
	}

	if _, ok := target.(Dummy); !ok {
		t.Error("Type assertion failed!")
	}

	sss := target.(Dummy)

	t.Error(sss.Field1)
	t.Error(sss.Field2.Foo())
}
