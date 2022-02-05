package generator

import (
	"errors"
	"fmt"

	schema "github.com/lestrrat-go/jsschema"
)

type StructConfigMap map[string]string

func (s StructConfigMap) Get(object *schema.Schema) (string, error) {
	name := s[object.Title]
	if name == "" {
		return "", fmt.Errorf("no struct config for title %s", object.Title)
	}
	return name, nil
}

func (s StructConfigMap) Set(name string, object *schema.Schema) {
	s[object.Title] = golangifyName(name)
}

func BuildPropertyConfig(name string, object *schema.Schema, scm StructConfigMap) (PropertyConfig, error) {
	var t Type
	var refStructGoName string
	var err error

	if len(object.Type) == 1 {
		switch object.Type[0] {
		case schema.ObjectType:
			t = ObjectType
			refStructGoName, err = scm.Get(object)
			if err != nil {
				return PropertyConfig{}, err
			}
		case schema.ArrayType:
			if len(object.Items.Schemas) != 1 {
				return PropertyConfig{}, errors.New("invalid array type")
			}

			if len(object.Items.Schemas[0].Type) != 1 {
				t = ArrayType
				refStructGoName = golangifyFieldRef(object.Items.Schemas[0].Reference)
			} else {
				switch object.Items.Schemas[0].Type[0] {
				case schema.StringType:
					t = ArrayStringType
				case schema.IntegerType:
					t = ArrayIntegerType
				case schema.NumberType:
					t = ArrayNumberType
				case schema.BooleanType:
					t = BoolType
				case schema.ObjectType:
					t = ArrayType
					refStructGoName, err = scm.Get(object.Items.Schemas[0])
					if err != nil {
						return PropertyConfig{}, err
					}
				default:
					return PropertyConfig{}, fmt.Errorf("invalid property sub-type for array %d", object.Items.Schemas[0].Type[0])
				}
			}
		case schema.StringType:
			t = StringType
		case schema.IntegerType:
			t = IntegerType
		case schema.NumberType:
			t = NumberType
		case schema.BooleanType:
			t = BoolType
		default:
			return PropertyConfig{}, fmt.Errorf("unknown property type %d", object.Type[0])
		}
	} else {
		t = ObjectType
		refStructGoName = golangifyFieldRef(object.Reference)
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

func BuildStructConfig(root *schema.Schema) ([]StructConfig, error) {
	if root == nil {
		return nil, errors.New("schema cannot be nil")
	}

	scm := StructConfigMap{}
	for name, object := range root.Definitions {
		if len(object.Type) != 1 || object.Type[0] != schema.ObjectType {
			continue
		}
		scm.Set(name, object)
	}

	var structs []StructConfig
	for name, object := range root.Definitions {
		if len(object.Type) != 1 || object.Type[0] != schema.ObjectType {
			continue
		}

		name = stripPrefix(name)

		var props []PropertyConfig

		for propName, propObject := range object.Properties {
			prop, err := BuildPropertyConfig(propName, propObject, scm)
			if err != nil {
				continue
				// return nil, fmt.Errorf("failed to build property config for %s: %w", propName, err)
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
