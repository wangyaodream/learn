package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	// err = client.Ping(ctx, nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB")

	collection := client.Database("school").Collection("student")

	s1 := Student{"Alice", 18}
	s2 := Student{"Bob", 31}
	s3 := Student{"Carter", 24}

	students := []interface{}{s1, s2, s3}

	insertResult, err := collection.InsertMany(ctx, students)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert result: ", insertResult)

	// TODO CRUD

}
