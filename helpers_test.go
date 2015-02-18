package digo

type Dummy struct {
	Field1 string
	Field2 DummyInterface
}

type Dummy2 struct {
}

func (d Dummy2) Foo() string {
	return "From Dummy2"
}

type DummyInterface interface {
	Foo() string
}
