package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type RestaurantModel struct {
	ID           primitive.ObjectID `bson:"_id"`
	RestaurantId string             `bson:"restaurant_id"`
	DishIds      []string           `bson:"dish_ids"`
	Name         string             `bson:"name"`
	Locations    []LocationModel    `bson:"locations,omitempty"`
	Types        []string           `bson:"types"`
}
