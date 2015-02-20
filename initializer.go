package digo

type Initializer interface {
	BeforeInject() error
}
