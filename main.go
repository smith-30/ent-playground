package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bxcodec/faker"
	"github.com/smith-30/ent-playground/ent"
	"github.com/smith-30/ent-playground/ent/car"
	"github.com/smith-30/ent-playground/ent/user"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type FakeData struct {
	Name string `faker:"name"`
}

func main() {
	client, err := ent.Open("mysql", "root:ent@tcp(localhost:13307)/ent_sample?parseTime=true")
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

	fuu, err := CreateCars(ctx, client, fa.Name)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("%#v\n", fuu)

	if err := QueryCars(ctx, fuu); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
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

func CreateCars(ctx context.Context, client *ent.Client, userName string) (*ent.User, error) {
	// creating new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}

	// creating new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}
	log.Println("car was created: ", ford)

	// create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("QueryCars().All failed querying user cars: %v", err)
	}
	log.Println("returned cars:", cars)

	// what about filtering specific cars.
	ford, err := a8m.QueryCars().
		Where(car.ModelEQ("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("QueryCars failed querying user cars: %v", err)
	}
	log.Println(ford)
	return nil
}

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	// query the inverse edge.
	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %v", ca.Model, err)
		}
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}
	return nil
}
