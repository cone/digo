package digo

import (
	"testing"
)

func TestContext_Unmarshal(t *testing.T) {
	path := "test-data/test.json"

	ctx := new(Context)

	err := ctx.unmarshal(path)
	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	if len(ctx.NodeMap) < 1 {
		t.Error("Incorrect number of nodes")
		return
	}

	if ctx.NodeMap["kitchen"].Type != "digo.Kitchen" {
		t.Error("Incorrect type")
	}
}

func TestContext_Get(t *testing.T) {
	initTypeRegistry()

	path := "test-data/test.json"

	ctx := new(Context)

	err := ctx.unmarshal(path)
	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	kitchenInterface, err := ctx.Get("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	if _, ok := kitchenInterface.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	kitchen := kitchenInterface.(Kitchen)

	if kitchen.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if kitchen.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	//pointer types are singletons by default
	//so this must modify the fridgr.temp of all
	//kitchen type objects
	kitchen.MyFridge.SetTemp(100)

	//a second kitchen

	kitchenInterface, _ = ctx.Get("kitchen")

	kitchen2 := kitchenInterface.(Kitchen)

	if kitchen2.MyFridge.GetTemp() != -1 {
		t.Error("Fridge is a copy by default so temp should be -1")
	}

	clearTypeRegistry()
}
