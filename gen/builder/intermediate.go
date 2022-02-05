package builder

import (
	"github.com/nikitawootten/oscal-neo4j/gen/config"
)

type IntermediateStruct struct {
	Title       string
	Description string
	Properties  []*IntermediateProperty
}

func (is IntermediateStruct) Key() string {
	return StructKey(is.Title, is.Description)
}

type IntermediateProperty struct {
	Name         string
	Title        string
	Description  string
	Type         config.Type
	StructRelKey string
}
