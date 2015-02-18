package digo

import (
	"reflect"
)

type DependencyNode struct {
	Name         string
	Type         reflect.Type
	FieldName    string
	Dependencies []*DependencyNode
}
