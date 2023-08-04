package surveyor_restaurant_test

import (
	"testing"

	mock_surveyor_restaurant "github.com/allergeye/surveyor-service/pkg/surveyor/mocks/restaurant"
	. "github.com/allergeye/surveyor-service/pkg/surveyor/restaurant"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func newMarshallersMock(t *testing.T) *mock_surveyor_restaurant.MockMarshallers {
	ctrl := gomock.NewController(t)
	return mock_surveyor_restaurant.NewMockMarshallers(ctrl)
}

func newFakeMarhsallers(mm *mock_surveyor_restaurant.MockMarshallers) MarshallersImplementation {
	return MarshallersImplementation{}
}

func Test_Marshaller_MarshalRestaurantRequestBody(t *testing.T) {
	restaurantRequest := AddRestaurantRequestBody{
		Name: "Restaurant1",
		Locations: []AddRestaurantLocationRequestBody{
			{
				StreetAddressLine1: "Restaurant1 Street",
				StreetAddressLine2: "",
				City:               "City",
				Province:           "Province",
				Country:            "Country",
				PostalCode:         "PostalCode",
			},
			{
				StreetAddressLine1: "Restaurant2 Street",
				StreetAddressLine2: "",
				City:               "City",
				Province:           "Province",
				Country:            "Country",
				PostalCode:         "PostalCode",
			},
		},
		Dishes: []AddRestaurantDishRequestBody{
			{
				Name: "Dish 1",
				Allergens: []AddRestaurantDishAllergenRequestBody{
					{
						Name:        "SESAME",
						Probability: 100,
					},
					{
						Name:        "PEANUT",
						Probability: 100,
					},
				},
			},
			{
				Name: "Dish 2",
				Allergens: []AddRestaurantDishAllergenRequestBody{
					{
						Name:        "EGGS",
						Probability: 100,
					},
					{
						Name:        "MILK",
						Probability: 100,
					},
				},
			},
		},
	}

	tests := map[string]struct {
		mocks       func() *mock_surveyor_restaurant.MockMarshallers
		expectedErr error
	}{
		"successfully marshalls a well-formed restaurant request body": {
			mocks: func() *mock_surveyor_restaurant.MockMarshallers {
				mm := newMarshallersMock(t)
				return mm
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mm := tt.mocks()
			m := newFakeMarhsallers(mm)

			_, _, err := m.MarshalRestaurantRequestBody(restaurantRequest)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
