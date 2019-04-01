package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type User struct {
	Email   string    `bson:"email" json:"email"`
	Password   string    `bson:"last_piece" json:"last_piece"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("testing").Collection("users")
	user := User{"hjbello", "1234"}
	collection.InsertOne(ctx, user)

	userOut := User{}
	err := collection.FindOne(ctx, bson.M{"email":"hj2bello"}).Decode(&userOut)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userOut)


}