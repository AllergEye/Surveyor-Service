package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type DishModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	DishId    string             `bson:"dish_id"`
	Name      string             `bson:"names"`
	Allergens []AllergenModel    `bson:"allergens,omitempty"`
}
