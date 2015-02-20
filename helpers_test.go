package digo

type Kitchen struct {
	Msg      string
	MyFridge Fridge
	MyStove  Stove
}

type Fridge interface {
	Freeze() string
	SetTemp(int)
	GetTemp() int
}

type Stove interface {
	Fry() string
}

type SuperFridge struct {
	temp int
}

func (this *SuperFridge) Freeze() string {
	return "Super Freeze"
}

func (this *SuperFridge) SetTemp(degrees int) {
	this.temp = degrees
}

func (this *SuperFridge) GetTemp() int {
	return this.temp
}

type OldStove struct{}

func (this OldStove) Fry() string {
	return "Frying slooooowly"
}
