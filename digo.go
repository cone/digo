package digo

import (
	"errors"
)

var Digo *ContextManager

func init() {
	if Digo == nil {
		Digo = new(ContextManager)
		Digo.contexts = map[string]*Context{}
	}
}

type ContextManager struct {
	contexts map[string]*Context
}

func (this *ContextManager) Context(filePath string) (*Context, error) {
	if ctx, exists := this.contexts[filePath]; exists {
		return ctx, nil
	}

	ctx, err := this.newContext(filePath)
	if err != nil {
		return nil, err
	}

	this.contexts[filePath] = ctx

	return ctx, nil

}

func (this *ContextManager) newContext(filePath string) (*Context, error) {
	ctx := &Context{}

	err := ctx.Unmarshal(filePath)
	if err != nil {
		return nil, errors.New("Error creating new Context -> " + err.Error())
	}

	return ctx, nil
}
