package digo

import (
	"reflect"
)

type DependencyNode struct {
	TypeName     string            `json:"type"`
	Type         reflect.Type      `json: "-"`
	FieldName    string            `json:"field"`
	Dependencies []*DependencyNode `json:"deps"`
}
