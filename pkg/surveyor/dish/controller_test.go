package surveyor_dish_test

import (
	"context"
	"errors"
	"testing"

	"github.com/allergeye/surveyor-service/internal/domain/dish"
	. "github.com/allergeye/surveyor-service/pkg/surveyor/dish"
	mock_surveyor_dish "github.com/allergeye/surveyor-service/pkg/surveyor/mocks/dish"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type controllerMock struct {
	logger      *zap.SugaredLogger
	dishService *mock_surveyor_dish.MockDishService
	marhsallers *mock_surveyor_dish.MockMarshallers
}

func newControllerMock(t *testing.T) controllerMock {
	ctrl := gomock.NewController(t)
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	dishservice := mock_surveyor_dish.NewMockDishService(ctrl)
	marshallers := mock_surveyor_dish.NewMockMarshallers(ctrl)
	return controllerMock{
		logger:      sugar,
		dishService: dishservice,
		marhsallers: marshallers,
	}
}

func newFakeController(cm controllerMock) DishControllerImplementation {
	return DishControllerImplementation{
		Logger:      cm.logger,
		DishService: cm.dishService,
		Marshallers: cm.marhsallers,
	}
}

func Test_Controller_NewDishController(t *testing.T) {
	t.Run("returns a new controller with expected values", func(t *testing.T) {
		cm := newControllerMock(t)

		expectedController := DishControllerImplementation{
			Logger:      cm.logger,
			DishService: cm.dishService,
			Marshallers: cm.marhsallers,
		}

		newController := NewDishContoller(cm.logger, cm.dishService, cm.marhsallers)
		assert.Equal(t, expectedController, newController)
	})
}

func Test_Controller_AddDishesToRestaurant(t *testing.T) {
	requestBody := AddDishesToRestaurantRequestBody{
		RestaurantId: "restuarant1",
		Dishes: []AddDishesToRestaurantDishesRequestBody{
			{
				Name: "dish 1",
				Allergens: []AddDishesToRestaurantDishesAllergensRequestBody{
					{
						Name:               "SESAME",
						IsProbabilityKnown: true,
						Probability:        100,
					},
				},
			},
			{
				Name:      "dish 2",
				Allergens: []AddDishesToRestaurantDishesAllergensRequestBody{},
			},
		},
	}

	restaurantId := "restaurant1"

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
		mocks       func() controllerMock
		expectedErr error
	}{
		"successfully adds dishes to restaurant": {
			mocks: func() controllerMock {
				cm := newControllerMock(t)
				cm.marhsallers.EXPECT().MarshalAddDishesToRestaurantRequestBody(requestBody).Return(restaurantId, dishes, nil)
				cm.dishService.EXPECT().AddDishesToRestaurant(gomock.Any(), restaurantId, dishes).Return(nil)
				return cm
			},
		},
		"returns an error if the request body could not be marshalled": {
			mocks: func() controllerMock {
				cm := newControllerMock(t)
				cm.marhsallers.EXPECT().MarshalAddDishesToRestaurantRequestBody(requestBody).Return("", []dish.Dish{}, randomErr)
				return cm
			},
			expectedErr: randomErr,
		},
		"returns an error if the dishes could not be added to the restaurant": {
			mocks: func() controllerMock {
				cm := newControllerMock(t)
				cm.marhsallers.EXPECT().MarshalAddDishesToRestaurantRequestBody(requestBody).Return(restaurantId, dishes, nil)
				cm.dishService.EXPECT().AddDishesToRestaurant(gomock.Any(), restaurantId, dishes).Return(randomErr)
				return cm
			},
			expectedErr: randomErr,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cm := tt.mocks()
			c := newFakeController(cm)
			ctx := context.Background()

			err := c.AddDishesToRestaurant(ctx, requestBody)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
