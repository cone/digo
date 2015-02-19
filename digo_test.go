package digo

import (
	"testing"
)

func TestDigo_Context(t *testing.T) {
	TypeRegistry.Add(Kitchen{})
	TypeRegistry.Add(&SuperFridge{})
	TypeRegistry.Add(OldStove{})

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
