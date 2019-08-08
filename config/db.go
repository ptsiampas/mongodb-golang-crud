package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var BooksDB *mongo.Database

func init() {
	var err error
	url := "mongodb://root:1234@localhost"
	ctx, _:= context.WithTimeout(context.Background(), 5*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalln("url for database is incorrect,", err)
	}

	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("unable to ping the database, make sure its running and authentication is correct, err:",err)
	}

	log.Println("connected to database", url)

	BooksDB = c.Database("books") // Set the database for use throughout the site.

}
