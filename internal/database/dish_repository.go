package database

import (
	"context"
	"fmt"

	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"go.mongodb.org/mongo-driver/mongo"
)

type DishRepository interface {
	AddDishes(ctx context.Context, dishes []dish.Dish) error
}

type DishRepositoryImplementation struct {
	client *mongo.Client
}

func NewDishRepository(client *mongo.Client) DishRepository {
	return DishRepositoryImplementation{
		client: client,
	}
}

func (dr DishRepositoryImplementation) AddDishes(ctx context.Context, dishes []dish.Dish) error {
	models := make([]interface{}, len(dishes))
	for i, dish := range dishes {
		d, err := marshalDish(dish)
		if err != nil {
			return fmt.Errorf("dishRepository.AddDishes: %w, %v", err, err)
		}
		models[i] = *d
	}

	coll := dr.client.Database("allergeye").Collection("dishes")
	_, err := coll.InsertMany(ctx, models)
	if err != nil {
		return err
	}
	return nil
}
