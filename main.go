package main

import (
	"context"
	"log"

	"github.com/smith-30/ent-playground/ent"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:ent@tcp(localhost:13307)/ent_sample")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
