package digo

var Digo *ContextManager

func init() {
	if Digo == nil {
		Digo = new(ContextManager)
		Digo.contexts = map[string]*Context{}
	}
}
