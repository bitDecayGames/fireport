package pogo

// TypedMsg is a struct for embedding. Never should be instantiated by itself
type TypedMsg struct {
	Type string `json:"msgType"`
}

// Typer is a basic interface for assigning a type field to a struct
type Typer interface {
	SetType(typeName string)
}

// SetType will assign the given type
func (t *TypedMsg) SetType(typeName string) {
	t.Type = typeName
}
