package digo

import (
	"reflect"
	"testing"
)

func TestDigo_Context(t *testing.T) {
	//TODO: this is not working
	//TypeRegistry.Add(Kitchen{})
	//TypeRegistry.Add(&SuperFridge{})
	//TypeRegistry.Add(OldStove{})

	TypeRegistry["digo.Kitchen"] = reflect.TypeOf(Kitchen{})
	TypeRegistry["digo.SuperFridge"] = reflect.TypeOf(&SuperFridge{})
	TypeRegistry["digo.OldStove"] = reflect.TypeOf(OldStove{})

	ctx, err := Digo.Context("test-data/test.json")
	if err != nil {
		t.Error("An error has ocurred:", err)
		return
	}

	i, err := ctx.Get("kitchen")
	if err != nil {
		t.Error("An error has ocurred:", err)
		return
	}

	if _, ok := i.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		t.Error(i)
	}
}
