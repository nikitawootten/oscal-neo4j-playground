package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/mindstand/gogm/v2"
	"github.com/nikitawootten/oscal-neo4j/schema"
)

func saveOSCAL() error {
	raw, err := os.ReadFile("vault/NIST_SP-800-53_rev5_catalog.json")
	if err != nil {
		return err
	}

	var catalogRoot struct {
		Catalog schema.CatalogCatalog `json:"catalog"`
	}
	err = json.Unmarshal(raw, &catalogRoot)
	if err != nil {
		return err
	}

	sess, err := gogm.G().NewSessionV2(gogm.SessionConfig{
		AccessMode: gogm.AccessModeWrite,
	})
	if err != nil {
		return err
	}
	defer sess.Close()

	err = sess.SaveDepth(context.Background(), catalogRoot.Catalog, 6)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	config := gogm.Config{
		IndexStrategy: gogm.VALIDATE_INDEX,
		PoolSize:      50,
		Port:          7687,
		IsCluster:     false,
		Host:          "0.0.0.0",
		Password:      "password",
		Username:      "neo4j",
	}
	// todo pass in list of nodes
	_gogm, err := gogm.New(&config, gogm.DefaultPrimaryKeyStrategy)
	if err != nil {
		panic(err)
	}
	gogm.SetGlobalGogm(_gogm)

	saveOSCAL()
}
