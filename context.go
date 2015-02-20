package digo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type ContextData struct {
	NodeMap map[string]*NodeData `json:"nodes"`
}

type Context struct {
	cache map[string]interface{}
	Nodes *ContextData
}

func (this *Context) Unmarshal(filePath string) error {
	data, err := this.getFileBytes(filePath)
	if err != nil {
		return errors.New("Error getting file data -> " + err.Error())
	}

	ctxData := &ContextData{}

	err = json.Unmarshal(data, ctxData)
	if err != nil {
		return errors.New("Error unmarshaling data -> " + err.Error())
	}

	this.Nodes = ctxData

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

func (this *Context) Get(key string) (interface{}, error) {
	var node *NodeData

	if tmpNode, exists := this.Nodes.NodeMap[key]; exists {
		node = tmpNode
	} else {
		return struct{}{}, errors.New("The given type cannot be found: " + key + " (forgot to add to the TypeRegister?)")
	}

	return this.memoize(key, node)
}

func (this *Context) memoize(key string, node *NodeData) (interface{}, error) {
	if cached, exists := this.cache[key]; exists {
		return cached, nil
	}

	t, err := depInjector.Resolve(node, this.Nodes.NodeMap)
	if err != nil {
		return t, err
	}

	this.cache[key] = t

	return t, nil
}

//TODO: Make an Maker interface so if the
//dependency to inject impements it , we can call
//it to initialize the dep

//TODO:See how to make singletons

//TODO: When an item in cache uses pointers, it shares it
//to all instances, check if we want that
