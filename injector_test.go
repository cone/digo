package digo

import (
	"reflect"
	"testing"
)

func TestInjector_New(t *testing.T) {
	injector := new(Injector)

	TypeRegistry["digo.Kitchen"] = reflect.TypeOf(Kitchen{})

	cp, err := injector.New("digo.Kitchen")
	if err != nil {
		t.Error("Type not found")
	}

	if _, ok := cp.(Kitchen); !ok {
		t.Error("Type assertion failed!")
	}
}

func TestInjector_Resolve(t *testing.T) {
	injector := new(Injector)

	TypeRegistry["digo.Kitchen"] = reflect.TypeOf(Kitchen{})
	TypeRegistry["digo.SuperFridge"] = reflect.TypeOf(&SuperFridge{})
	TypeRegistry["digo.OldStove"] = reflect.TypeOf(OldStove{})

	dependencyTree := &DependencyNode{
		Name: "digo.Kitchen",
		Dependencies: []*DependencyNode{
			{
				Name:      "digo.SuperFridge",
				FieldName: "MyFridge",
			},
		},
	}

	target, err := injector.Resolve(dependencyTree)
	if err != nil {
		t.Error("The error has ocurred", err)
	}

	if _, ok := target.(Kitchen); !ok {
		t.Error("Type assertion failed!")
	}

	asserted := target.(Kitchen)

	if asserted.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect value")
	}
}
