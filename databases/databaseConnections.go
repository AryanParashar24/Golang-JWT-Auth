package database

import (
	"fmt"
	"log"
	"time"
	"os"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	err:= godotenv.Load("env") // this will load the env file from the root directory of our project
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err:=mongo.NewCLient(option.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection established successfully")
	return client

}

var Client *mongo.client = DBinstance()	// variable client which is a type mongo.client whchi is a DBinstance funtion and then later it would be called to return 
// we call this function so that it can get returned along with the client which will be captured in the variable client

// Now we'll use a function to access a aprticular collection  of our DB
func OpenCollectionclient *mongo.Client, collectionName string) *mongo.Collection {	// so we can pass a collection here and then can use it 
// and then mongo.Collection will be the returned from the function OpenCollection
	var collection *mongo.Collection = client.Database("").Collection(collectionName)	// we can define the database in the client.Database and the name of the collection will be in the .Collection which has been named as the collectionName here 
	return collection	// will return the collection which will be used in the controllers to access the data from the database
}