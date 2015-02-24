package digo

import (
	"testing"
)

func TestDigo_Context(t *testing.T) {
	TypeRegistry.Add(Kitchen{})
	TypeRegistry.Add(SuperFridge{})
	TypeRegistry.Add(OldStove{})

	ctx, err := ContextFor("test-data/test.json")
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

func TestDigo_New(t *testing.T) {
	TypeRegistry.Add(Kitchen{})

	cp, err := digo.New("digo.Kitchen", false)
	if err != nil {
		t.Error("Type not found")
	}

	if _, ok := cp.(Kitchen); !ok {
		t.Error("Type assertion failed!")
	}

	TypeRegistry = TypeMap{}
}
