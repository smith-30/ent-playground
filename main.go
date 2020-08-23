package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bxcodec/faker"
	"github.com/smith-30/ent-playground/ent"
	"github.com/smith-30/ent-playground/ent/user"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type FakeData struct {
	Name string `faker:"name"`
}

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

	fa := FakeData{}

	if err := faker.FakeData(&fa); err != nil {
		fmt.Println(err)
		return
	}

	u, err := CreateUser(ctx, client, fa.Name)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("%#v\n", u)

	fu, err := QueryUser(ctx, client, fa.Name)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("%#v\n", fu)
}

func CreateUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.NameEQ(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
