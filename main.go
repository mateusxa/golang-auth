package main

import (
	"fmt"
	"log"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var ctx = context.TODO()

func init() {
	// credential := options.Credential{
	// 	Username: "adminuser",
	// 	Password: "password123",
	// }
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	userCollection = client.Database("auth").Collection("user")
}

func main() {
    // create user
	// read user
	// update user
	// delete user
	// authenticate user
	fmt.Println("Hello, world.")
	user := &User{
		ID:        	primitive.NewObjectID(),
		Name:      	"Mateus",
		Email: 		"mateus@mateus.com",
		Updated: 	time.Now(),
		Created: 	time.Now(),
	}

	err := createUser(user)
	if err != nil {
		fmt.Printf(err.Error())
		return 
	}
	fmt.Printf("user %s created", user.Name)
	return
}

type User struct {
	ID        	primitive.ObjectID 		`bson:"_id"`
	Name		string					`bson:"name"`
	Email		string					`bson:"email"`
	Password	string					`bson:"password"`
	Updated		time.Time				`bson:"updated"`
	Created		time.Time				`bson:"created"`
}

func createUser(user *User) error {
		_, err := userCollection.InsertOne(ctx, user)
	return err
}
