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

	//// GET

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

	//// Copy

	copiedInterface, err := ctx.Copy("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	if _, ok := copiedInterface.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	copied := copiedInterface.(Kitchen)

	if copied.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if copied.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if copied.Msg != "" {
		t.Error("Msg should be empty!")
	}

	if copied.MyFridge.GetTemp() != 0 {
		t.Error("Fridge is not shared so it should be 0")
	}

	//// Single

	singleInterface, err := ctx.Single("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	if _, ok := singleInterface.(*Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	single := singleInterface.(*Kitchen)

	if single.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if single.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if single.Msg != "" {
		t.Error("Msg should be empty!")
	}

	if single.MyFridge.GetTemp() != 0 {
		t.Error("Fridge temp should be 0")
	}

	single.MyFridge.SetTemp(200)

	anotherSI, err := ctx.Single("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	another := anotherSI.(*Kitchen)

	if another.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if another.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if another.Msg != "" {
		t.Error("Msg should be empty!")
	}

	if another.MyFridge.GetTemp() != 200 {
		t.Error("'another' is a singleton, so temp should be 200")
	}

	TypeRegistry = TypeMap{}
}
