package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/nikitawootten/oscal-neo4j/gen/builder"
	"github.com/nikitawootten/oscal-neo4j/gen/jsonschema"
	"github.com/nikitawootten/oscal-neo4j/gen/outputer"
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

	b := builder.NewBuilder()
	structs, err := b.Build(&s)
	if err != nil {
		panic(err)
	}

	outputer.Output(os.Stdout, "out", structs...)
}
