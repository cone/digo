package digo

import (
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
	TypeRegistry.Add(Kitchen{})
	TypeRegistry.Add(SuperFridge{})
	TypeRegistry.Add(OldStove{})

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

	kitchen.Msg = "kitchen"
	kitchen.MyFridge.SetTemp(10)

	//// a copy from cache

	i2, err := ctx.Get("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	if _, ok := i2.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	kitchen2 := i.(Kitchen)

	if kitchen2.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if kitchen2.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if kitchen2.Msg != "" {
		t.Error("Msg should be empty!")
	}

	if kitchen2.MyFridge.GetTemp() != 10 {
		t.Error("Fridge is shared (a pointer) so it temp sholud be 10")
	}

	TypeRegistry = TypeMap{}
}
