package surveyor_test

import (
	"context"
	"errors"
	"testing"

	"github.com/allergeye/surveyor-service/internal/domain/restaurant"
	mock_lib "github.com/allergeye/surveyor-service/pkg/surveyor/mocks/lib"
	mock_surveyor "github.com/allergeye/surveyor-service/pkg/surveyor/mocks/restaurant"
	. "github.com/allergeye/surveyor-service/pkg/surveyor/restaurant"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type controllerMock struct {
	logger            *zap.SugaredLogger
	restaurantService *mock_surveyor.MockRestaurantService
	helpers           *mock_lib.MockHelpers
}

func newControllerMock(t *testing.T) controllerMock {
	ctrl := gomock.NewController(t)
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	restaurantService := mock_surveyor.NewMockRestaurantService(ctrl)
	helpers := mock_lib.NewMockHelpers(ctrl)
	return controllerMock{
		logger:            sugar,
		restaurantService: restaurantService,
		helpers:           helpers,
	}
}

func newFakeController(cm controllerMock) RestaurantControllerImplementation {
	return RestaurantControllerImplementation{
		Logger:            cm.logger,
		RestaurantService: cm.restaurantService,
		Helpers:           cm.helpers,
	}
}

func Test_Controller_GetAllRestaurants(t *testing.T) {
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
		mocks       func() controllerMock
		expectedErr error
	}{
		"successfully gets all restaurants from the service": {
			mocks: func() controllerMock {
				cm := newControllerMock(t)
				cm.restaurantService.EXPECT().GetAllRestaurants(gomock.Any()).Return(restaurants, nil)
				return cm
			},
		},
		"returns an error if there was an error getting all the restaurants from the service": {
			mocks: func() controllerMock {
				cm := newControllerMock(t)
				cm.restaurantService.EXPECT().GetAllRestaurants(gomock.Any()).Return(nil, randomErr)
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

			_, err := c.GetAllRestaurants(ctx)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Controller_AddRestaurant(t *testing.T) {
	expectedRestaurant := restaurant.Restaurant{
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
	}

	randomErr := errors.New("random error")

	tests := map[string]struct {
		mocks       func() controllerMock
		expectedErr error
	}{
		"successfully adds a restaurant": {
			mocks: func() controllerMock {
				cm := newControllerMock(t)
				cm.helpers.EXPECT().GenerateUUID().Return(uuid.MustParse("52fdfc07-2182-454f-963f-5f0f9a621d72"))
				cm.restaurantService.EXPECT().AddRestaurant(gomock.Any(), expectedRestaurant).Return(nil)
				return cm
			},
		},
		"returns an error if the restaurant could not be added": {
			mocks: func() controllerMock {
				cm := newControllerMock(t)
				cm.helpers.EXPECT().GenerateUUID().Return(uuid.MustParse("52fdfc07-2182-454f-963f-5f0f9a621d72"))
				cm.restaurantService.EXPECT().AddRestaurant(gomock.Any(), expectedRestaurant).Return(randomErr)
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
			restaurantRequest := AddRestaurantRequestBody{
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
			}

			err := c.AddRestaurant(ctx, restaurantRequest)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
