package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Draft() {
	UserDB := db.Collection("users")

	type User struct {
		Id        primitive.ObjectID `bson:"_id"`
		Age       int                `bson:"age"`
		FirstName string             `bson:"first_name"`
		LastName  string             `bson:"last_name"`
	}

	user := User{}
	UserDB.FindOne(ctx, map[string]string{}).Decode(&user)
	fmt.Printf("user: %v\n", user)
}
