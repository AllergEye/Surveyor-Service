package surveyor_dish_test

import (
	"testing"

	. "github.com/allergeye/surveyor-service/pkg/surveyor/dish"
	mock_database "github.com/allergeye/surveyor-service/pkg/surveyor/mocks/database"
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
