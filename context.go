package digo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Context struct {
	singletons map[string]interface{} `json:"-"`
	Nodes      []*DependencyNode      `json:"nodes"`
}

func (this *Context) Unmarshal(filePath string) error {
	data, err := this.getFileBytes(filePath)
	if err != nil {
		return errors.New("Error getting file data -> " + err.Error())
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		return errors.New("Error unmarshaling data -> " + err.Error())
	}

	return nil
}

func (this *Context) getFileBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Error opening file -> " + err.Error())
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("Error reading file -> " + err.Error())
	}

	return fileBytes, nil
}

func (this *Context) Get(key string) (interface{}, error) {
	var node *DependencyNode

	for _, firstChildNode := range this.Nodes {
		if firstChildNode.TypeName == key {
			node = firstChildNode
			break
		}
	}

	if node == nil {
		return struct{}{}, errors.New("The given type cannot be found (forgot to add to the TypeRegister?)")
	}

	return depInjector.Resolve(node)
}
