package digo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeMap_Add_AddingString(t *testing.T) {
	registry := TypeMap{}

	test := "Hello"
	err := registry.Add(test)

	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	tp := registry["string"]

	if toString(tp) != "string" {
		t.Errorf("Incorrect type: %s", toString(tp))
	}
}

func TestTypeMap_Add_AddingStruct(t *testing.T) {
	registry := TypeMap{}

	test := Kitchen{}
	err := registry.Add(test)

	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	tp := registry["digo.Kitchen"]

	if toString(tp) != "digo.Kitchen" {
		t.Errorf("Incorrect type: %s", toString(tp))
	}
}

func TestTypeMap_AddType(t *testing.T) {
	registry := TypeMap{}

	test := "Hello"

	tp := reflect.TypeOf(test)

	err := registry.AddType(tp)

	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	tp = registry["string"]

	if toString(tp) != "string" {
		t.Error("Incorrect type")
	}
}

func TestTypeMap_Get(t *testing.T) {
	registry := TypeMap{}

	test := "Hello"

	tp := reflect.TypeOf(test)

	registry["string"] = tp

	gottenTp, err := registry.Get("string")
	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	if toString(gottenTp) != "string" {
		t.Error("Incorrect type")
	}
}

func toString(param reflect.Type) string {
	return fmt.Sprintf("%v", param)
}
