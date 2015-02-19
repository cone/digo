package digo

type NodeData struct {
	Type string     `json:"type"`
	Deps []*DepData `json:"deps"`
}

type DepData struct {
	ID    string `json:"id"`
	Field string `json:"field"`
}
