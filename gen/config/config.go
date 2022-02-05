package config

type Type string

const (
	ObjectType       Type = "OBJECT"
	ArrayType        Type = "ARRAY"
	StringType       Type = "string"
	IntegerType      Type = "int"
	NumberType       Type = "float"
	BoolType         Type = "bool"
	ArrayStringType  Type = "[]string"
	ArrayIntegerType Type = "[]int"
	ArrayNumberType  Type = "[]float"
	ArrayBoolType    Type = "[]bool"
)

type Struct struct {
	GoName      string
	Description string
	Properties  []Property
	Refs        []Reference
}

type Property struct {
	Name            string
	GoName          string
	Description     string
	Type            Type
	RefStructGoName string
}

type Reference struct {
	PropertyGoName string
	StructGoName   string
}
