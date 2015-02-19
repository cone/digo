package digo

import (
	"reflect"
	"testing"
)

func TestInjector_New(t *testing.T) {
	TypeRegistry["digo.Kitchen"] = reflect.TypeOf(Kitchen{})

	cp, err := depInjector.New("digo.Kitchen")
	if err != nil {
		t.Error("Type not found")
	}

	if _, ok := cp.(Kitchen); !ok {
		t.Error("Type assertion failed!")
	}

	TypeRegistry = TypeMap{}
}

func TestInjector_Resolve(t *testing.T) {
	TypeRegistry["digo.Kitchen"] = reflect.TypeOf(Kitchen{})
	TypeRegistry["digo.SuperFridge"] = reflect.TypeOf(&SuperFridge{})
	TypeRegistry["digo.OldStove"] = reflect.TypeOf(OldStove{})

	dependencyTree := &DependencyNode{
		TypeName: "digo.Kitchen",
		Dependencies: []*DependencyNode{
			{
				TypeName:  "digo.SuperFridge",
				FieldName: "MyFridge",
			},
		},
	}

	target, err := depInjector.Resolve(dependencyTree)
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

	TypeRegistry = TypeMap{}
}
