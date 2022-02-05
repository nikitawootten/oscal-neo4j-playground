package main

import (
	"fmt"
	"os"

	schema "github.com/lestrrat-go/jsschema"
	"github.com/nikitawootten/oscal-neo4j/gen/generator"
)

func main() {
	s, err := schema.ReadFile("../vault/oscal_schema.json")
	if err != nil {
		panic(fmt.Errorf("did you run prepare_schema.sh? %w", err))
	}

	structs, err := generator.BuildStructConfig(s)
	if err != nil {
		panic(err)
	}

	generator.Output(os.Stdout, "out", structs...)
}
