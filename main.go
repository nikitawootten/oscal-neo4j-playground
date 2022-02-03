package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nikitawootten/oscal-neo4j/schema"
)

func main() {
	raw, err := os.ReadFile("oscal-content/nist.gov/SP800-53/rev5/json/NIST_SP-800-53_rev5_catalog.json")
	if err != nil {
		log.Fatal(err)
	}

	var catalogRoot schema.Root
	err = json.Unmarshal(raw, &catalogRoot)
	if err != nil {
		log.Fatal(err)
	}

}
