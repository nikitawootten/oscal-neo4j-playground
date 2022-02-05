package generator

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

type StructConfig struct {
	Name        string
	GoName      string
	Description string
	Properties  []PropertyConfig
	Refs        []ReferenceConfig
}

type PropertyConfig struct {
	Name            string
	GoName          string
	Description     string
	Type            Type
	RefStructGoName string
}

type ReferenceConfig struct {
	PropertyGoName string
	StructGoName   string
}
