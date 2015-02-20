package digo

import (
	"testing"
)

func TestInjector_Resolve(t *testing.T) {
	TypeRegistry.Add(Kitchen{})
	TypeRegistry.Add(SuperFridge{})
	TypeRegistry.Add(OldStove{})

	ctxMap := map[string]*NodeData{
		"super_fridge": &NodeData{
			Type:  "digo.SuperFridge",
			IsPtr: true,
		},
		"kitchen": &NodeData{
			Type: "digo.Kitchen",
			Deps: []*DepData{
				&DepData{
					ID:    "super_fridge",
					Field: "MyFridge",
				},
			},
		},
	}

	target, err := depInjector.resolve(ctxMap["kitchen"], ctxMap)
	if err != nil {
		t.Error("The error has ocurred:", err)
		return
	}

	if _, ok := target.(Kitchen); !ok {
		t.Error("Type assertion failed!")
		return
	}

	asserted := target.(Kitchen)

	if asserted.MyFridge.Freeze() != "Super Freeze" {
		t.Error("Incorrect value")
	}

	TypeRegistry = TypeMap{}
}
