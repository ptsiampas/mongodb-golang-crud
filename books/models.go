package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongodb-web-dev/config"
	"time"
)

type Book struct {
	Id     primitive.ObjectID `bson:"_id"`
	Isbn   string             `bson:"isbn"`
	Title  string             `bson:"title"`
	Author string             `bson:"author"`
	Price  float64            `bson:"price"`
}

var Books *mongo.Collection

func init() {
	Books = config.BooksDB.Collection("books")

}

func AllBooks() []Book {
	var books []Book
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := Books.Find(ctx, bson.D{})
	if err != nil {
		log.Fatalln("allBooks:", err)
	}
	defer c.Close(ctx)
	x := 0
	for c.Next(ctx) {
		var b Book

		err := c.Decode(&b)
		if err != nil {
			log.Fatalln("allbooks-loop", err)
		}
		books = append(books, b)
		x++
	}
	return books
}

func UpdateBook(b Book) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	f := bson.M{"_id": b.Id}

	// I use ReplaceOne here because there is an underlying function within the driver that will check what has changed
	// and Update only that field. (coll.updateOrReplace)
	x, err := Books.ReplaceOne(ctx, f, b)

	if err != nil {
		log.Println("UpdateBook:", err)
		return "", err
	}

	log.Println("updated", x.ModifiedCount, "record with id", b.Id.Hex())
	return b.Isbn, err
}

func FindOneBook(s string) (Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"isbn": s}

	var b Book
	if err := Books.FindOne(ctx, filter).Decode(&b); err != nil {
		return Book{}, err
	}

	return b, nil
}

func AddBook(b Book) (string, error) {
	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Because its a new structure the id is initialised with zeros, so we need to assign a real ObjectId here.
	b.Id = primitive.NewObjectID()

	oid, err := Books.InsertOne(ctx, b)
	if err != nil {
		log.Println(err)
		return "", err
	}

	isbn, err := getBookIsbnById(oid.InsertedID.(primitive.ObjectID))

	return isbn, err
}

// DeleteBook Takes an ISBN and returns the number of delete books. Which should normally be 1, but you know mongo.
func DeleteBook(s string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := Books.DeleteOne(ctx, bson.M{"isbn": s})
	if err != nil {
		return 0, err
	}
	return int(res.DeletedCount), nil
}

// getBookIsbnById will get a mongo Object Id and return the isbn for that record.
func getBookIsbnById(s primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": s}
	opt := options.FindOne().SetProjection(bson.M{"isbn": 1})

	var b Book
	if err := Books.FindOne(ctx, filter, opt).Decode(&b); err != nil {
		log.Println("getBookIsbnById:", err)
		return "", err
	}

	return b.Isbn, nil
}
