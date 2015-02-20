package digo

var TypeRegistry TypeMap

func init() {
	if TypeRegistry == nil {
		TypeRegistry = TypeMap{}
	}
}
