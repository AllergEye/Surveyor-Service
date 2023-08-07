package dish

import "github.com/google/uuid"

type Dish struct {
	DishId    string
	Name      string
	Allergens []Allergen
}

func NewDish(name string, allergens []Allergen) Dish {
	return Dish{
		DishId:    uuid.NewString(),
		Name:      name,
		Allergens: allergens,
	}
}
