package databaseTests

import (
	mongodbdal "CMPSC488SP24SecThursday/bam"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

func TestDatabaseOperation(t *testing.T) {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongodbdal.MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection was successful
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Println("Successfully connected to MongoDB!")
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			t.Error(err)
		}
	}(client, context.Background())

	// Get a handle for the collection
	collection := client.Database("test").Collection("users")

	// Inserting a user
	hashedPassword, err := mongodbdal.HashPassword("secret_password")
	if err != nil {
		log.Fatal(err)
	}

	user1 := mongodbdal.User{
		Username: "exampleUser1",
		Password: hashedPassword,
	}

	user2 := mongodbdal.User{
		Username: "exampleUser2",
		Password: hashedPassword,
	}

	fmt.Println("User insertion")
	fmt.Println("User to be inserted:", user1)
	fmt.Println("User to be inserted:", user2)
	fmt.Println()

	insertResult1, err := collection.InsertOne(context.Background(), user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult1.InsertedID)

	insertResult2, err := collection.InsertOne(context.Background(), user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult2.InsertedID)

	// Print the entire contents of the database
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for cursor.Next(context.Background()) {
		var user mongodbdal.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Println("User:", user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// Finding a user
	var result mongodbdal.User
	filter := mongodbdal.User{Username: "exampleUser1", Password: hashedPassword}
	err = collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found user:", result)

	// Comparing password
	err = mongodbdal.ComparePasswords(result.Password, "secret_password")
	if err != nil {
		fmt.Println("Password does not match")
	} else {
		fmt.Println("Password matches")
	}

	// Dropping the entire database
	err = client.Database("test").Drop(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database dropped successfully.")
}
