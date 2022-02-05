package jsonschema

type Root struct {
	Comment     string                 `json:"$comment"`
	ID          string                 `json:"$id"`
	Schema      string                 `json:"$schema"`
	Type        string                 `json:"type"`
	Definitions map[string]*Definition `json:"definitions"`
}

type Definition struct {
	ID          string                 `json:"$id"`
	Ref         string                 `json:"$ref"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Type        string                 `json:"type"`
	Items       *Definition            `json:"items"`
	Properties  map[string]*Definition `json:"properties"`
}
