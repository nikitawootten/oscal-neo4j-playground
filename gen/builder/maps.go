package builder

import (
	"fmt"

	"github.com/nikitawootten/oscal-neo4j/gen/jsonschema"
)

type Ref struct {
	Name string
	Ref  *jsonschema.Definition
}

// Track JSON refs
type RefMap map[string]*Ref

func (r RefMap) Get(ref string) *Ref {
	return r[ref]
}

func (r *RefMap) Set(name string, object *jsonschema.Definition) {
	(*r)[object.ID] = &Ref{
		Name: name,
		Ref:  object,
	}
}

// Track structs
type StructMap map[string]*IntermediateStruct

func (s StructMap) Get(title, description string) *IntermediateStruct {
	return s[StructKey(title, description)]
}

func (s *StructMap) Set(is *IntermediateStruct) {
	(*s)[is.Key()] = is
}

// Track struct names
type NameMap map[string]string

func (n NameMap) Get(is *IntermediateStruct) string {
	return n[is.Key()]
}

func (n *NameMap) Set(name string, is *IntermediateStruct) {
	(*n)[is.Key()] = name
}

type VisitedMap map[string]bool

func (v *VisitedMap) Visited(title, description string) bool {
	key := StructKey(title, description)
	_, ok := (*v)[key]
	(*v)[key] = true
	return ok
}

type RelatedMap map[string]map[string]map[string]bool

func (r *RelatedMap) Relate(childKey, parentKey, propName string) {
	_, ok := (*r)[childKey]
	if !ok {
		(*r)[childKey] = map[string]map[string]bool{parentKey: {propName: true}}
	} else {
		_, ok = (*r)[childKey][parentKey]
		if !ok {
			(*r)[childKey][parentKey] = map[string]bool{propName: true}
		} else {
			(*r)[childKey][parentKey][propName] = true
		}
	}
}

func StructKey(title, description string) string {
	return fmt.Sprintf("%s::%s", title, description)
}
