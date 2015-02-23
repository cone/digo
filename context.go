package digo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Context struct {
	cache   map[string]interface{} `json:"-"`
	NodeMap map[string]*NodeData   `json:"nodes"`
}

func (this *Context) unmarshal(filePath string) error {
	data, err := this.getFileBytes(filePath)
	if err != nil {
		return errors.New("Error getting file data -> " + err.Error())
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		return errors.New("Error unmarshaling data -> " + err.Error())
	}

	this.cache = map[string]interface{}{}

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

func (this *Context) memoize(key string, node *NodeData) (interface{}, error) {
	if cached, exists := this.cache[key]; exists {
		return cached, nil
	}

	t, err := depInjector.resolve(node, key, this.NodeMap)
	if err != nil {
		return t, err
	}

	this.cache[key] = t

	return t, nil
}

func (this *Context) Get(key string) (interface{}, error) {
	node, err := this.getFromNodeMap(key)
	if err != nil {
		return struct{}{}, err
	}

	return depInjector.resolve(node, key, this.NodeMap)
}

func (this *Context) getFromNodeMap(key string) (*NodeData, error) {
	var node *NodeData

	if tmpNode, exists := this.NodeMap[key]; exists {
		node = tmpNode
	} else {
		return nil, errors.New("The given type cannot be found: " + key + " (forgot to add to the TypeRegister?)")
	}

	return node, nil
}
