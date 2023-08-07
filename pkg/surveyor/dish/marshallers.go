package surveyor_dish

import "github.com/allergeye/surveyor-service/internal/domain/dish"

type Marshallers interface {
	MarshalAddDishesToRestaurantRequestBody(addDishesToRestaurantRequestBody AddDishesToRestaurantRequestBody) (string, []dish.Dish, error)
	MarshalAddDishesToRestaurantDishesRequestBody(addDishesToRestaurantDishesRequestBody []AddDishesToRestaurantDishesRequestBody) ([]dish.Dish, error)
	MarshalAddDishesToRestaurantDishesAllergensRequestBody(addDishesToRestaurantDishesAllergensRequestBody []AddDishesToRestaurantDishesAllergensRequestBody) ([]dish.Allergen, error)
}

type MarshallersImplementation struct{}

func NewMarshallers() Marshallers {
	return MarshallersImplementation{}
}

func (m MarshallersImplementation) MarshalAddDishesToRestaurantRequestBody(addDishesToRestaurantRequestBody AddDishesToRestaurantRequestBody) (string, []dish.Dish, error) {
	restaurantId := addDishesToRestaurantRequestBody.RestaurantId

	dishes, err := m.MarshalAddDishesToRestaurantDishesRequestBody(addDishesToRestaurantRequestBody.Dishes)
	if err != nil {
		return "", []dish.Dish{}, err
	}

	return restaurantId, dishes, nil
}

func (m MarshallersImplementation) MarshalAddDishesToRestaurantDishesRequestBody(addDishesToRestaurantDishesRequestBody []AddDishesToRestaurantDishesRequestBody) ([]dish.Dish, error) {
	dishes := make([]dish.Dish, len(addDishesToRestaurantDishesRequestBody))
	for i, requestDish := range addDishesToRestaurantDishesRequestBody {
		allergens, err := m.MarshalAddDishesToRestaurantDishesAllergensRequestBody(requestDish.Allergens)
		if err != nil {
			return []dish.Dish{}, err
		}
		dish := dish.NewDish(requestDish.Name, allergens)
		dishes[i] = dish
	}

	return dishes, nil
}

func (m MarshallersImplementation) MarshalAddDishesToRestaurantDishesAllergensRequestBody(addDishesToRestaurantDishesAllergensRequestBody []AddDishesToRestaurantDishesAllergensRequestBody) ([]dish.Allergen, error) {
	allergens := make([]dish.Allergen, len(addDishesToRestaurantDishesAllergensRequestBody))
	for i, requestAllergen := range addDishesToRestaurantDishesAllergensRequestBody {
		if !dish.IsValidAllergenName(requestAllergen.Name) {
			return []dish.Allergen{}, ErrInvalidAllergen
		}

		if (requestAllergen.IsProbabilityKnown && requestAllergen.Probability != 100) || (!requestAllergen.IsProbabilityKnown && requestAllergen.Probability != 0) {
			return []dish.Allergen{}, ErrUserCannotGuessProbability
		}

		allergen := dish.NewAllergen(requestAllergen.Name, requestAllergen.IsProbabilityKnown, requestAllergen.Probability)
		allergens[i] = allergen
	}

	return allergens, nil
}
