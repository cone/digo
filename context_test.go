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

	//Modifying non shared field
	kitchen.Msg = "kitchen"

	//Modifying shared field (pointer)
	kitchen.MyFridge.SetTemp(10)

	otherKitchenInterface, err := ctx.Get("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	if _, ok := otherKitchenInterface.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	otherKitchen := otherKitchenInterface.(Kitchen)

	if otherKitchen.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if otherKitchen.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if otherKitchen.Msg != "" {
		t.Error("Msg should be empty!")
	}

	if otherKitchen.MyFridge.GetTemp() != 10 {
		t.Error("Fridge is shared (a pointer) so it temp sholud be 10")
	}

	clearTypeRegistry()
}

func TestContext_Copy(t *testing.T) {
	initTypeRegistry()

	path := "test-data/test.json"

	ctx := new(Context)

	err := ctx.unmarshal(path)
	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	copiedInterface, err := ctx.Copy("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	anotherInterface, err := ctx.Copy("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	if _, ok := copiedInterface.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	copied := copiedInterface.(Kitchen)
	another := anotherInterface.(Kitchen)

	if copied.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if copied.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if copied.Msg != "" {
		t.Error("Msg should be empty!")
	}

	if copied.MyFridge.GetTemp() != -1 {
		t.Error("Fridge is not shared so it should be -1")
	}

	copied.Msg = "Hello"
	copied.MyFridge.SetTemp(20)

	//The other copy should hold its own values

	if another.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if another.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if another.Msg != "" {
		t.Error("Msg should be empty!")
	}

	if another.MyFridge.GetTemp() != -1 {
		t.Error("Fridge is not shared so it should be -1")
	}

	clearTypeRegistry()
}

func TestContext_Single(t *testing.T) {
	initTypeRegistry()

	path := "test-data/test.json"

	ctx := new(Context)

	err := ctx.unmarshal(path)
	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	singleInterface, err := ctx.Single("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	anotherInterface, err := ctx.Single("kitchen")
	if err != nil {
		t.Error("An error has ocurred: ", err)
		return
	}

	if _, ok := singleInterface.(*Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	single := singleInterface.(*Kitchen)
	another := anotherInterface.(*Kitchen)

	if single.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if single.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if single.Msg != "" {
		t.Error("Msg should be empty!")
	}

	if single.MyFridge.GetTemp() != -1 {
		t.Error("Fridge is not shared so it should be -1")
	}

	single.Msg = "Hello"
	single.MyFridge.SetTemp(20)

	//another should be the same

	if another.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect Output")
	}

	if another.MyStove.Fry() != "Frying slooooowly" {
		t.Error("Incorrect Output")
	}

	if another.Msg != "Hello" {
		t.Error("Msg should be the same!")
	}

	if another.MyFridge.GetTemp() != 20 {
		t.Error("The fridge temp should be the same")
	}

	clearTypeRegistry()
}
