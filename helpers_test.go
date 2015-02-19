package digo

type Kitchen struct {
	MyFridge Fridge
	MyStove  Stove
}

type Fridge interface {
	Freeze() string
}

type Stove interface {
	Fry() string
}

type SuperFridge struct{}

func (this *SuperFridge) Freeze() string {
	return "Super Freeze"
}

type OldStove struct{}

func (this OldStove) Fry() string {
	return "Frying slooooowly"
}
