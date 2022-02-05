package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/nikitawootten/oscal-neo4j/gen/generator"
	"github.com/nikitawootten/oscal-neo4j/gen/jsonschema"
)

func main() {
	raw, err := os.Open("../vault/oscal_schema.json")
	if err != nil {
		panic(err)
	}
	defer raw.Close()

	rawBytes, err := io.ReadAll(raw)
	if err != nil {
		panic(err)
	}

	var s jsonschema.Root
	json.Unmarshal(rawBytes, &s)

	structs, err := generator.BuildStructConfig(&s)
	if err != nil {
		panic(err)
	}

	generator.Output(os.Stdout, "out", structs...)
}
