package surveyor_dish_test

import (
	"context"
	"errors"
	"testing"

	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	. "github.com/allergeye/surveyor-service/pkg/surveyor/dish"
	mock_database "github.com/allergeye/surveyor-service/pkg/surveyor/mocks/database"
	"github.com/stretchr/testify/assert"
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

func newFakeService(sm serviceMock) DishServiceImplementation {
	return DishServiceImplementation{
		Logger:         sm.logger,
		RestaurantRepo: sm.restaurantRepo,
		DishRepo:       sm.dishRepo,
	}
}

func Test_Service_NewDishService(t *testing.T) {
	t.Run("returns a new service with the expected values", func(t *testing.T) {
		sm := newServiceMock(t)

		expectedService := DishServiceImplementation{
			Logger:         sm.logger,
			RestaurantRepo: sm.restaurantRepo,
			DishRepo:       sm.dishRepo,
		}

		newService := NewDishService(sm.logger, sm.restaurantRepo, sm.dishRepo)
		assert.Equal(t, expectedService, newService)
	})
}

func Test_Service_GetDishesByRestaurantId(t *testing.T) {
	restaurantId := "restaurant-id"

	dishes := []dish.Dish{
		{
			DishId: "dish1",
			Name:   "Dish 1",
			Allergens: []dish.Allergen{
				{
					Name:               "SESAME",
					IsProbabilityKnown: true,
					Probability:        100,
				},
			},
		},
		{
			DishId:    "dish2",
			Name:      "Dish 2",
			Allergens: []dish.Allergen{},
		},
	}

	restaurant := restaurant.Restaurant{
		RestaurantId: "restaurant1",
		Name:         "Restaurant 1",
		DishIds:      []string{"dish1", "dish2"},
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() serviceMock
		expectedErr error
	}{
		"successfully gets a restaurant's dishes by its id": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetRestaurantById(gomock.Any(), restaurantId).Return(&restaurant, nil)
				sm.dishRepo.EXPECT().GetDishById(gomock.Any(), restaurant.DishIds[0]).Return(&dishes[0], nil)
				sm.dishRepo.EXPECT().GetDishById(gomock.Any(), restaurant.DishIds[1]).Return(&dishes[1], nil)
				return sm
			},
		},
		"returns an error if the restaurant could not be retrieved by id": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetRestaurantById(gomock.Any(), restaurantId).Return(nil, randomErr)
				return sm
			},
			expectedErr: randomErr,
		},
		"returns an error if a given dish could not be retrieved by its id": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().GetRestaurantById(gomock.Any(), restaurantId).Return(&restaurant, nil)
				sm.dishRepo.EXPECT().GetDishById(gomock.Any(), restaurant.DishIds[0]).Return(nil, randomErr)
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

			_, err := s.GetDishesByRestaurantId(ctx, restaurantId)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Service_AddDishesToRestaurant(t *testing.T) {
	restaurantId := "restaurant1"

	dishIds := []string{"dish1", "dish2"}

	dishes := []dish.Dish{
		{
			DishId: "dish1",
			Name:   "dish 1",
			Allergens: []dish.Allergen{
				{
					Name:               "SESAME",
					IsProbabilityKnown: true,
					Probability:        100,
				},
			},
		},
		{
			DishId:    "dish2",
			Name:      "dish 2",
			Allergens: []dish.Allergen{},
		},
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() serviceMock
		expectedErr error
	}{
		"successfully adds dishes to a given restaurant": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().AddDishesToRestaurant(gomock.Any(), restaurantId, dishIds).Return(nil)
				sm.dishRepo.EXPECT().AddDishes(gomock.Any(), dishes).Return(nil)
				return sm
			},
		},
		"returns an error if the given dishes could not be added to a restaurant": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().AddDishesToRestaurant(gomock.Any(), restaurantId, dishIds).Return(randomErr)
				return sm
			},
			expectedErr: randomErr,
		},
		"returns an error if the dishes could not be added": {
			mocks: func() serviceMock {
				sm := newServiceMock(t)
				sm.restaurantRepo.EXPECT().AddDishesToRestaurant(gomock.Any(), restaurantId, dishIds).Return(nil)
				sm.dishRepo.EXPECT().AddDishes(gomock.Any(), dishes).Return(randomErr)
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

			err := s.AddDishesToRestaurant(ctx, restaurantId, dishes)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
