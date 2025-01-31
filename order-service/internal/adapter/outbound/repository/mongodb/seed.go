package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedDatabase(client *mongo.Client, dbName string) error {
	customerCollection := client.Database(dbName).Collection("customers")
	restaurantCollection := client.Database(dbName).Collection("restaurants")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := customerCollection.InsertMany(ctx, bson.A{
		bson.M{"_id": "d215b5f8-0249-4dc5-89a3-51fd148cfb41", "user_name": "user_1", "first_name": "First", "last_name": "User"},
		bson.M{"_id": "d215b5f8-0249-4dc5-89a3-51fd148cfb42", "user_name": "user_2", "first_name": "Second", "last_name": "User"},
	})

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Print("Seed data already exists")
		}
		return err
	}

	documents := bson.A{
		bson.M{
			"_id":    "d215b5f8-0249-4dc5-89a3-51fd148cfb45",
			"name":   "restaurant_1",
			"active": true,
			"products": []bson.M{
				{
					"_id":    "d215b5f8-0249-4dc5-89a3-51fd148cfb47",
					"name":   "product_1",
					"price":  25.0,
					"active": false},
				{
					"_id":    "d215b5f8-0249-4dc5-89a3-51fd148cfb48",
					"name":   "product_2",
					"price":  50.0,
					"active": true},
			}},
		bson.M{
			"_id":    "d215b5f8-0249-4dc5-89a3-51fd148cfb46",
			"name":   "restaurant_2",
			"active": false,
			"products": []bson.M{
				{
					"_id":    "d215b5f8-0249-4dc5-89a3-51fd148cfb49",
					"name":   "product_3",
					"price":  20.00,
					"active": false},
				{
					"_id":    "d215b5f8-0249-4dc5-89a3-51fd148cfb50",
					"name":   "product_4",
					"price":  40.00,
					"active": true},
			}},
	}

	_, err = restaurantCollection.InsertMany(ctx, documents)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Print("Seed data already exists")
		}
		return err
	}

	log.Print("Seed data inserted successfully")
	return nil
}
