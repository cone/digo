package digo

type NodeData struct {
	Type        string     `json:"type"`
	IsPtr       bool       `json:"is_pointer"`
	IsSingleton bool       `json:"is_singleton"`
	Deps        []*DepData `json:"deps"`
}

type DepData struct {
	ID    string `json:"id"`
	Field string `json:"field"`
}
