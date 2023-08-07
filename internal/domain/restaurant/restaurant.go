package restaurant

import "github.com/google/uuid"

type Restaurant struct {
	RestaurantId string
	DishIds      []string
	Name         string
	Locations    []Location
}

func NewRestaurant(dishIds []string, name string, locations []Location) Restaurant {
	return Restaurant{
		RestaurantId: uuid.NewString(),
		DishIds:      dishIds,
		Name:         name,
		Locations:    locations,
	}
}
