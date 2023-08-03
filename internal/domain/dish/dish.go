package dish

import "github.com/google/uuid"

type Dish struct {
	DishId    uuid.UUID
	Name      string
	Allergens []Allergen
}

func NewDish(name string, allergens []Allergen) Dish {
	return Dish{
		DishId:    uuid.New(),
		Name:      name,
		Allergens: allergens,
	}
}
