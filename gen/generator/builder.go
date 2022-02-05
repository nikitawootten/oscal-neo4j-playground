package generator

import (
	"errors"
	"fmt"

	"github.com/nikitawootten/oscal-neo4j/gen/jsonschema"
)

type RefMap map[string]*Ref

func (r RefMap) Get(ref string) *Ref {
	return r[ref]
}

func (r RefMap) Set(name string, object *jsonschema.Definition) {
	r[object.ID] = &Ref{
		Name: name,
		Ref:  object,
	}
}

type Ref struct {
	Name string
	Ref  *jsonschema.Definition
}

func BuildPropertyConfig(name string, object *jsonschema.Definition, refMap RefMap) (PropertyConfig, error) {
	var t Type
	var refStructGoName string

	if object.Ref != "" {
		ref := refMap.Get(object.Ref)
		if ref == nil {
			return PropertyConfig{}, errors.New("unknown reference")
		}

		object = ref.Ref
	}

	switch object.Type {
	case "object":
		// special case
	case "array":
		if len(object.Items.Type) != 1 {
			t = ArrayType
			refStructGoName = golangifyFieldRef(object.Items.Ref)
		} else {
			subObj := object.Items
			if subObj.Ref != "" {
				ref := refMap.Get(object.Ref)
				if ref == nil {
					return PropertyConfig{}, errors.New("unknown reference")
				}

				subObj = ref.Ref
			}

			if len(subObj.Type) != 1 {
				return PropertyConfig{}, errors.New("no subtype specified")
			}

			switch subObj.Type {
			case "string":
				t = ArrayStringType
			case "integer":
				t = ArrayIntegerType
			case "number":
				t = ArrayNumberType
			case "boolean":
				t = BoolType
			case "object":
				// special case
			default:
				return PropertyConfig{}, fmt.Errorf("invalid property sub-type for array %s", object.Items.Type)
			}
		}
	case "string":
		t = StringType
	case "integer":
		t = IntegerType
	case "number":
		t = NumberType
	case "boolean":
		t = BoolType
	default:
		return PropertyConfig{}, fmt.Errorf("unknown property type %d", object.Type[0])
	}

	config := PropertyConfig{
		Name:            name,
		GoName:          golangifyName(name),
		Description:     object.Description,
		Type:            t,
		RefStructGoName: refStructGoName,
	}

	return config, nil
}

func BuildStructConfig(root *jsonschema.Root) ([]StructConfig, error) {
	if root == nil {
		return nil, errors.New("schema cannot be nil")
	}

	refMap := RefMap{}
	for name, object := range root.Definitions {
		if object.ID != "" {
			refMap.Set(name, object)
		}
	}

	var structs []StructConfig
	for name, object := range root.Definitions {
		if object.Type != "object" {
			continue
		}

		name = stripPrefix(name)

		var props []PropertyConfig

		for propName, propObject := range object.Properties {
			prop, err := BuildPropertyConfig(propName, propObject, refMap)
			if err != nil {
				return nil, fmt.Errorf("failed to build property config for %s: %w", propName, err)
			}

			props = append(props, prop)
		}

		s := StructConfig{
			Name:       name,
			GoName:     golangifyName(name),
			Properties: props,
		}
		structs = append(structs, s)
	}

	return structs, nil
}
