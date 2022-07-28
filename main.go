package main

import (
	"context"
	"fmt"

	commonMongo "gitlab.deeporium.com/apps/common/adapter/mongodb"
)

func main() {
	database, _ := commonMongo.New(&commonMongo.Config{
		WriteURL: "mongodb://localhost:27017",
		ReadURL:  "mongodb://localhost:27017",
		Database: "company",
	})

	repository := NewRepository(database)
	r, _ := Repository[Customer](*repository).FindById(context.Background(), "62e0e653116111b3c0c3f220")
	r2, _ := Repository[Inspector](*repository).FindById(context.Background(), "62e0e653116111b3c0c3f220")
	fmt.Println(r)
	fmt.Println(r2)

}
