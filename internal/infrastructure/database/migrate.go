package database

import (
	"context"
	"fmt"
	"log"
	"time"

	// "github.com/google/uuid"
	// "go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Migrate(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.CreateCollection(ctx, "spatial")
	if err != nil {
		// if !mongo.IsDuplicateKeyError(err) {
		// 	log.Println("Collection already exist")
		// }
		log.Println("Collection already exist")
	}

	collection := db.Collection("spatial")
	collection = db.Collection("test")

	indexModel := mongo.IndexModel{
		Keys: map[string]interface{}{
			"id": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	// res, insertErr := collection.InsertOne(ctx, bson.M{
	// 	// "id":   uuid.New().String(),
	// 	"name": "test",
	// 	// "email": fmt.Sprintf("test_%d@example.com", time.Now().Unix()),
	// })
	// if insertErr != nil {
	// 	fmt.Println("Insert error", insertErr)
	// 	return insertErr
	// }

	// fmt.Println("Inserted ID", res.InsertedID)
	fmt.Println("DB", db.Name())
	fmt.Println("Collection", collection.Name())

	fmt.Println("Migrated Complete")
	// count, _ := collection.CountDocuments(ctx, bson.M{})
	// fmt.Println("📊 Total docs:", count)

	return nil
}
