package digo

import (
	"reflect"
	"testing"
)

func TestContext_Unmarshal(t *testing.T) {
	path := "test-data/test.json"

	ctx := new(Context)

	err := ctx.Unmarshal(path)
	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	if len(ctx.Nodes.NodeMap) < 1 {
		t.Error("Incorrect number of nodes")
		return
	}

	if ctx.Nodes.NodeMap["kitchen"].Type != "digo.Kitchen" {
		t.Error("Incorrect type")
	}
}

func TestContext_Get(t *testing.T) {
	TypeRegistry["digo.Kitchen"] = reflect.TypeOf(Kitchen{})
	TypeRegistry["digo.SuperFridge"] = reflect.TypeOf(&SuperFridge{})
	TypeRegistry["digo.OldStove"] = reflect.TypeOf(OldStove{})

	path := "test-data/test.json"

	ctx := new(Context)

	err := ctx.Unmarshal(path)
	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	i, err := ctx.Get("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	if _, ok := i.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	kitchen := i.(Kitchen)

	if kitchen.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if kitchen.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	TypeRegistry = TypeMap{}
}
