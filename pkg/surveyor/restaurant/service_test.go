package surveyor_restaurant_test

import (
	"context"
	"errors"
	"testing"

	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	mock_database "github.com/allergeye/surveyor-service/pkg/surveyor/mocks/database"
	. "github.com/allergeye/surveyor-service/pkg/surveyor/restaurant"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type serviceMock struct {
	logger         *zap.SugaredLogger
	restaurantRepo *mock_database.MockRestaurantRepository
	dishRepo       *mock_database.MockDishRepository
}

func newServiceMock(t *testing.T) serviceMock {
	ctrl := gomock.NewController(t)
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	restaurantRepo := mock_database.NewMockRestaurantRepository(ctrl)
	dishRepo := mock_database.NewMockDishRepository(ctrl)

	return serviceMock{
		logger:         sugar,
		restaurantRepo: restaurantRepo,
		dishRepo:       dishRepo,
	}
}

func newFakeService(sm serviceMock) RestaurantServiceImplementation {
	return RestaurantServiceImplementation{
		Logger:         sm.logger,
		RestaurantRepo: sm.restaurantRepo,
		DishRepo:       sm.dishRepo,
	}
}

func Test_Service_GetAllRestaurants(t *testing.T) {
	restaurants := []restaurant.Restaurant{
		{
			Name: "Restaurant1",
			Locations: []restaurant.Location{
				{
					StreetAddressLine1: "Restaurant1 Street",
					StreetAddressLine2: "",
					City:               "City",
					Province:           "Province",
					Country:            "Country",
					PostalCode:         "PostalCode",
				},
			},
		},
		{
			Name: "Restaurant2",
			Locations: []restaurant.Location{
				{
					StreetAddressLine1: "Restaurant2 Street",
					StreetAddressLine2: "",
					City:               "City",
					Province:           "Province",
					Country:            "Country",
					PostalCode:         "PostalCode",
				},
			},
		},
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() serviceMock
		expectedErr error
	}{
		"successfully returns all restaurants": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetAllRestaurants(gomock.Any()).Return(restaurants, nil)
				return sm
			},
		},
		"returns an error if there was an error getting all of the restaurants from the repository": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetAllRestaurants(gomock.Any()).Return(nil, randomErr)
				return sm
			},
			expectedErr: randomErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sm := tt.mocks()
			s := newFakeService(sm)
			ctx := context.Background()

			_, err := s.GetAllRestaurants(ctx)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Service_AddRestaurant(t *testing.T) {
	restaurantName := "Restaurant1"

	restaurantToAdd := restaurant.Restaurant{
		RestaurantId: uuid.MustParse("52fdfc07-2182-454f-963f-5f0f9a621d72"),
		Name:         "Restaurant1",
		Locations: []restaurant.Location{
			{
				StreetAddressLine1: "Restaurant1 Street",
				StreetAddressLine2: "",
				City:               "City",
				Province:           "Province",
				Country:            "Country",
				PostalCode:         "PostalCode",
			},
		},
		DishIds: []uuid.UUID{uuid.New(), uuid.New()},
	}

	newLocations := []restaurant.Location{
		{
			StreetAddressLine1: "Restaurant1 Street",
			StreetAddressLine2: "",
			City:               "City",
			Province:           "Province",
			Country:            "Country",
			PostalCode:         "PostalCode",
		},
		{
			StreetAddressLine1: "Restaurant1 Avenue",
			StreetAddressLine2: "",
			City:               "City",
			Province:           "Province",
			Country:            "Country",
			PostalCode:         "PostalCode",
		},
	}

	existingRestaurant1 := restaurant.Restaurant{
		RestaurantId: uuid.MustParse("52fdfc07-2182-454f-963f-5f0f9a621d72"),
		Name:         "Restaurant1",
		Locations: []restaurant.Location{
			{
				StreetAddressLine1: "Restaurant1 Road",
				StreetAddressLine2: "",
				City:               "City",
				Province:           "Province",
				Country:            "Country",
				PostalCode:         "PostalCode",
			},
		},
		DishIds: []uuid.UUID{uuid.New(), uuid.New()},
	}

	dishesToAdd := []dish.Dish{
		{
			Name: "Dish 1",
			Allergens: []dish.Allergen{
				{
					Name:        "SESAME",
					Probability: 100,
				},
				{
					Name:        "PEANUTS",
					Probability: 100,
				},
			},
		},
		{
			Name: "Dish 2",
			Allergens: []dish.Allergen{
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
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() serviceMock
		expectedErr error
	}{
		"successfully adds a new restaurant": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetRestaurantByName(gomock.Any(), restaurantName).Return(nil, mongo.ErrNoDocuments)
				sm.dishRepo.EXPECT().AddDishes(gomock.Any(), dishesToAdd).Return(nil)
				sm.restaurantRepo.EXPECT().AddRestaurant(gomock.Any(), restaurantToAdd).Return(nil)
				return sm
			},
		},
		"successfully adds a new restaurant location because a restaurant with the given name already exists": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetRestaurantByName(gomock.Any(), restaurantName).Return(&existingRestaurant1, nil)
				sm.restaurantRepo.EXPECT().UpdateRestaurantLocations(gomock.Any(), existingRestaurant1, []restaurant.Location{newLocations[0]}).Return(nil)
				return sm
			},
		},
		"returns an error if there was an error updating the restaurant's location": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetRestaurantByName(gomock.Any(), restaurantName).Return(&existingRestaurant1, nil)
				sm.restaurantRepo.EXPECT().UpdateRestaurantLocations(gomock.Any(), existingRestaurant1, []restaurant.Location{newLocations[0]}).Return(randomErr)
				return sm
			},
			expectedErr: randomErr,
		},
		"returns an ErrRestaurantAlreadyExists if a restaurant with the given name and locations already exists": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetRestaurantByName(gomock.Any(), restaurantName).Return(&restaurantToAdd, nil)
				return sm
			},
			expectedErr: ErrRestaurantAlreadyExists,
		},
		"returns an error if a new restaurant could not be added": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetRestaurantByName(gomock.Any(), restaurantName).Return(nil, mongo.ErrNoDocuments)
				sm.dishRepo.EXPECT().AddDishes(gomock.Any(), dishesToAdd).Return(nil)
				sm.restaurantRepo.EXPECT().AddRestaurant(gomock.Any(), restaurantToAdd).Return(randomErr)
				return sm
			},
			expectedErr: randomErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sm := tt.mocks()
			s := newFakeService(sm)
			ctx := context.Background()

			err := s.AddRestaurant(ctx, restaurantToAdd, dishesToAdd)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
