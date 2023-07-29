package database

import (
	"context"
	"fmt"

	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantRepository interface {
	GetAllRestaurants() ([]restaurant.Restaurant, error)
	GetRestaurantByName(name string) (*restaurant.Restaurant, error)
	AddRestaurant(restaurant restaurant.Restaurant) error
}

type RestaurantRepositoryImplementation struct {
	client *mongo.Client
}

func NewRestaurantRepository(client *mongo.Client) RestaurantRepository {
	return RestaurantRepositoryImplementation{
		client: client,
	}
}

func (r RestaurantRepositoryImplementation) GetAllRestaurants() ([]restaurant.Restaurant, error) {
	models := []*RestaurantModel{}
	coll := r.client.Database("allergeye").Collection("restaurants")

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf("restaurantRepository.GetAllRestaurants: %w: %v", err, err)
	}
	if err = cursor.All(context.TODO(), &models); err != nil {
		return nil, fmt.Errorf("restaurantRepository.GetAllRestaurants: %w: %v", err, err)
	}

	restaurants := make([]restaurant.Restaurant, len(models))

	for i, model := range models {
		unmarshalled, err := unmarshalRestaurant(model)
		if err != nil {
			return nil, fmt.Errorf("restaurantRepository.GetAllRestaurants: %w: %v", err, err)
		}
		restaurants[i] = *unmarshalled
	}

	return restaurants, nil
}

func (r RestaurantRepositoryImplementation) AddRestaurant(restaurant restaurant.Restaurant) error {
	model, err := marshalRestaurant(restaurant)
	if err != nil {
		return fmt.Errorf("restaurantRepository.AddRestaurant: %w, %v", err, err)
	}
	coll := r.client.Database("allergeye").Collection("restaurants")
	_, err = coll.InsertOne(context.TODO(), model)
	if err != nil {
		return fmt.Errorf("restaurantRepository.AddRestaurant: %w, %v", err, err)
	}

	return nil
}

func (r RestaurantRepositoryImplementation) GetRestaurantByName(name string) (*restaurant.Restaurant, error) {
	var result restaurant.Restaurant

	coll := r.client.Database("allergeye").Collection("restaurants")
	filter := bson.D{{"name", name}}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("restaurantRepository.GetRestaurantByName: %w, %v", err, err)
	}

	return &result, nil
}
