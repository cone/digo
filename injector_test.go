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

	ctxMap := map[string]*NodeData{
		"super_fridge": &NodeData{
			Type: "digo.SuperFridge",
		},
		"kitchen": &NodeData{
			Type: "digo.Kitchen",
			Deps: []*NodeData{
				&NodeData{
					ID:    "super_fridge",
					Field: "MyFridge",
				},
			},
		},
	}

	target, err := depInjector.Resolve(ctxMap["kitchen"], ctxMap)
	if err != nil {
		t.Error("The error has ocurred:", err)
		return
	}

	if _, ok := target.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	asserted := target.(Kitchen)

	if asserted.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect value")
	}

	TypeRegistry = TypeMap{}
}
