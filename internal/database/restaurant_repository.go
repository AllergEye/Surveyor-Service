package database

import (
	"context"
	"fmt"

	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RestaurantRepository interface {
	GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error)
	GetRestaurantByName(ctx context.Context, name string) (*restaurant.Restaurant, error)
	GetRestaurantById(ctx context.Context, restaurantId string) (*restaurant.Restaurant, error)
	AddRestaurant(ctx context.Context, restaurant restaurant.Restaurant) error
	AddDishesToRestaurant(ctx context.Context, restaurantId string, dishIds []string) error
	UpdateRestaurantLocations(ctx context.Context, restaurant restaurant.Restaurant, locations []restaurant.Location) error
}

type RestaurantRepositoryImplementation struct {
	client *mongo.Client
}

func NewRestaurantRepository(client *mongo.Client) RestaurantRepository {
	return RestaurantRepositoryImplementation{
		client: client,
	}
}

func (r RestaurantRepositoryImplementation) GetAllRestaurants(ctx context.Context) ([]restaurant.Restaurant, error) {
	models := []*RestaurantModel{}
	coll := r.client.Database("allergeye").Collection("restaurants")

	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf("restaurantRepository.GetAllRestaurants: %w: %v", err, err)
	}
	if err = cursor.All(ctx, &models); err != nil {
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

func (r RestaurantRepositoryImplementation) GetRestaurantByName(ctx context.Context, name string) (*restaurant.Restaurant, error) {
	var result RestaurantModel

	coll := r.client.Database("allergeye").Collection("restaurants")
	filter := bson.D{{"name", name}}

	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("restaurantRepository.GetRestaurantByName: %w, %v", err, err)
	}

	restaurant, err := unmarshalRestaurant(&result)
	if err != nil {
		return nil, fmt.Errorf("restaurantRepository.GetRestaurantByName: %w, %v", err, err)
	}

	return restaurant, nil
}

func (r RestaurantRepositoryImplementation) GetRestaurantById(ctx context.Context, restaurantId string) (*restaurant.Restaurant, error) {
	var result RestaurantModel

	coll := r.client.Database("allergeye").Collection("restaurants")
	filter := bson.D{{"restaurant_id", restaurantId}}

	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("restaurantRepository.GetRestaurantById: %w, %v", err, err)
	}

	restaurant, err := unmarshalRestaurant(&result)
	if err != nil {
		return nil, fmt.Errorf("restaurantRepository.GetRestaurantById: %w, %v", err, err)
	}

	return restaurant, nil
}

func (r RestaurantRepositoryImplementation) AddRestaurant(ctx context.Context, restaurant restaurant.Restaurant) error {
	model, err := marshalRestaurant(restaurant)
	if err != nil {
		return fmt.Errorf("restaurantRepository.AddRestaurant: %w, %v", err, err)
	}
	coll := r.client.Database("allergeye").Collection("restaurants")
	_, err = coll.InsertOne(ctx, model)
	if err != nil {
		return fmt.Errorf("restaurantRepository.AddRestaurant: %w, %v", err, err)
	}

	return nil
}

func (r RestaurantRepositoryImplementation) AddDishesToRestaurant(ctx context.Context, restaurantId string, dishIds []string) error {
	coll := r.client.Database("allergeye").Collection("restaurants")

	filter := bson.D{{"restaurant_id", restaurantId}}
	change := bson.M{"$push": bson.M{"dish_ids": bson.M{"$each": dishIds}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedRestaurant RestaurantModel
	err := coll.FindOneAndUpdate(ctx, filter, change, opts).Decode(&updatedRestaurant)
	if err != nil {
		return fmt.Errorf("restaurantRepository.AddDishesToRestaurant: %w, %v", err, err)
	}

	return nil
}

func (r RestaurantRepositoryImplementation) UpdateRestaurantLocations(ctx context.Context, restaurant restaurant.Restaurant, locations []restaurant.Location) error {
	models, err := marshalLocations(locations)
	if err != nil {
		return fmt.Errorf("restaurantRepository.UpdateRestaurantLocations: %w, %v", err, err)
	}

	coll := r.client.Database("allergeye").Collection("restaurants")
	filter := bson.D{{"restaurant_id", restaurant.RestaurantId}}
	change := bson.M{"$push": bson.M{"locations": bson.M{"$each": models}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedRestaurant RestaurantModel
	err = coll.FindOneAndUpdate(ctx, filter, change, opts).Decode(&updatedRestaurant)
	if err != nil {
		return fmt.Errorf("restaurantRepository.UpdateRestaurantLocations: %w, %v", err, err)
	}

	return nil
}
