package digo

var digo *ContextManager

func init() {
	if digo == nil {
		digo = new(ContextManager)
		digo.contexts = map[string]*Context{}
	}
}

func ContextFor(filePath string) (*Context, error) {
	return digo.Context(filePath)
}
