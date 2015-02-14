package digo

import (
	"fmt"
	"reflect"
	"testing"
)

type Dummy struct {
	Field1 string
}

func TestTypeMap_Add(t *testing.T) {
	resetTypeMap()

	test := "Hello"
	err := TypeRegistry.Add(test)

	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	tp := TypeRegistry["string"]

	if toString(tp) != "string" {
		t.Error("Incorrect type")
	}
}

func TestTypeMap_AddType(t *testing.T) {
	resetTypeMap()

	test := "Hello"

	tp := reflect.TypeOf(test)

	err := TypeRegistry.AddType(tp)

	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	tp := TypeRegistry["string"]

	if toString(tp) != "string" {
		t.Error("Incorrect type")
	}
}

func TestTypeMap_Get(t *testing.T) {
	resetTypeMap()

	tp := reflect.TypeOf(test)

	TypeRegistry["string"] = tp

	gottenTp, err := TypeRegistry.Get()
	if err != nil {
		t.Error("An error has ocurred: ", err)
	}

	if toString(gottenTp) != "string" {
		t.Error("Incorrect type")
	}
}

func resetTypeMap() {
	TypeRegistry = TypeMap{}
}

func toString(param reflect.Type) string {
	return fmt.Sprintf("%v", param)
}
