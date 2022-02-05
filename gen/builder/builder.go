package builder

import (
	"errors"
	"fmt"

	"github.com/nikitawootten/oscal-neo4j/gen/config"
	"github.com/nikitawootten/oscal-neo4j/gen/jsonschema"
)

type Builder struct {
	refMap     RefMap
	structMap  StructMap
	nameMap    NameMap
	visitedMap VisitedMap
}

func (b *Builder) processPropertyDefinition(name string, object *jsonschema.Definition) (*IntermediateProperty, error) {
	var t config.Type
	var structRelKey string

	if object.Ref != "" {
		ref := b.refMap.Get(object.Ref)
		if ref == nil {
			return nil, fmt.Errorf("unknown reference %s", object.Ref)
		}

		object = ref.Ref
	}

	switch object.Type {
	case "object":
		t = config.ObjectType

		if !b.visitedMap.Visited(object.Title, object.Description) {
			if _, err := b.processStructDefinition(object); err != nil {
				return nil, err
			}
		}
		structRelKey = StructKey(object.Title, object.Description)
	case "array":
		subObj := object.Items
		if subObj.Ref != "" {
			ref := b.refMap.Get(subObj.Ref)
			if ref == nil {
				return nil, fmt.Errorf("unknown reference %s", subObj.Ref)
			}

			subObj = ref.Ref
		}

		switch subObj.Type {
		case "string":
			t = config.ArrayStringType
		case "integer":
			t = config.ArrayIntegerType
		case "number":
			t = config.ArrayNumberType
		case "boolean":
			t = config.ArrayBoolType
		case "object":
			t = config.ArrayType

			if !b.visitedMap.Visited(subObj.Title, subObj.Description) {
				if _, err := b.processStructDefinition(subObj); err != nil {
					return nil, err
				}
			}
			structRelKey = StructKey(subObj.Title, subObj.Description)
		default:
			return nil, fmt.Errorf("invalid property sub-type for array %s", object.Items.Type)
		}
	case "string":
		t = config.StringType
	case "integer":
		t = config.IntegerType
	case "number":
		t = config.NumberType
	case "boolean":
		t = config.BoolType
	default:
		return nil, fmt.Errorf("unknown property type %d", object.Type[0])
	}

	return &IntermediateProperty{
		Name:         name,
		Title:        object.Title,
		Description:  object.Description,
		Type:         t,
		StructRelKey: structRelKey,
	}, nil
}

func (b *Builder) processStructDefinition(object *jsonschema.Definition) (*IntermediateStruct, error) {
	if object == nil {
		return nil, errors.New("definition cannot be nil")
	}

	is := b.structMap.Get(object.Title, object.Description)
	if is != nil {
		return is, nil
	}

	var props []*IntermediateProperty
	for propName, propObject := range object.Properties {
		prop, err := b.processPropertyDefinition(propName, propObject)
		if err != nil {
			return nil, fmt.Errorf("failed to build property config for %s: %w", propName, err)
		}

		props = append(props, prop)
	}

	is = &IntermediateStruct{
		Title:       object.Title,
		Description: object.Description,
		Properties:  props,
	}

	b.structMap.Set(is)

	return is, nil
}

func (b *Builder) Build(root *jsonschema.Root) ([]config.Struct, error) {
	if root == nil {
		return nil, errors.New("schema cannot be nil")
	}

	// build JSON ref map
	for name, object := range root.Definitions {
		if object.ID != "" {
			b.refMap.Set(name, object)
		}
	}

	// build intermediate struct definitions
	for name, object := range root.Definitions {
		if object.Type != "object" {
			continue
		}

		is, err := b.processStructDefinition(object)
		if err != nil {
			return nil, err
		}
		b.nameMap.Set(name, is)
	}

	// finalize name struct
	for key, intermediateStruct := range b.structMap {
		name, ok := b.nameMap[key]
		if !ok {
			name = intermediateStruct.Title
		}
		b.nameMap.Set(golangifyName(stripPrefix(name)), intermediateStruct)
	}

	// build final struct definitions
	structs := []config.Struct{}
	for key, intermediateStruct := range b.structMap {
		name, ok := b.nameMap[key]
		if !ok {
			return nil, errors.New("unknown name")
		}

		properties := []config.Property{}
		for _, intermediateProperty := range intermediateStruct.Properties {
			refStructGoName := ""
			if intermediateProperty.StructRelKey != "" {
				refStructGoName = b.nameMap[intermediateProperty.StructRelKey]
				if refStructGoName == "" {
					return nil, errors.New("unknown name")
				}
			}

			properties = append(properties, config.Property{
				Name:            intermediateProperty.Name,
				GoName:          golangifyName(stripPrefix(intermediateProperty.Name)),
				Description:     intermediateProperty.Description,
				Type:            intermediateProperty.Type,
				RefStructGoName: refStructGoName,
			})
		}

		structs = append(structs, config.Struct{
			GoName:      name,
			Description: intermediateStruct.Description,
			Properties:  properties,
		})
	}

	return structs, nil
}

func NewBuilder() *Builder {
	return &Builder{
		refMap:     RefMap{},
		structMap:  StructMap{},
		nameMap:    NameMap{},
		visitedMap: VisitedMap{},
	}
}
