package repository

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantRepository struct {
	client               *mongo.Client
	restaurantCollection *mongo.Collection
}

func NewRestaurantRepository(client *mongo.Client, dbName string) *RestaurantRepository {
	return &RestaurantRepository{
		client:               client,
		restaurantCollection: client.Database(dbName).Collection("restaurants"),
	}
}

func (r *RestaurantRepository) FindRestaurant(ctx context.Context,
	restaurantID string,
	productIds []string) (*entity.Restaurant, error) {
	var restaurantDocument RestaurantDocument
	criteria := bson.A{}
	for _, productId := range productIds {
		criteria = append(criteria, bson.M{
			"$elemMatch": bson.M{"product_id": productId},
		})
	}

	filter := bson.M{"restaurant_id": restaurantID,
		"products": bson.M{"$all": criteria}}

	err := r.restaurantCollection.FindOne(ctx, filter).Decode(&restaurantDocument)
	if err != nil {
		return nil, err
	}
	restaurant := ToRestaurantDomainEntity(&restaurantDocument)
	return restaurant, nil
}
