package surveyor_dish

import (
	"errors"

	"github.com/allergeye/surveyor-service/internal/domain/dish"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ErrInvalidRestaurantId = errors.New("invalid restaurant id")
)

type DishController interface {
	GetDishesByRestaurantId(c *gin.Context) ([]dish.Dish, error)
	AddDishesToRestaurant(c *gin.Context, requestBody AddDishesToRestaurantRequestBody) error
}

type DishControllerImplementation struct {
	Logger      *zap.SugaredLogger
	DishService DishService
	Marshallers Marshallers
}

func NewDishContoller(logger *zap.SugaredLogger, dishService DishService, marshallers Marshallers) DishController {
	return DishControllerImplementation{
		Logger:      logger,
		DishService: dishService,
		Marshallers: marshallers,
	}
}

func (dc DishControllerImplementation) GetDishesByRestaurantId(c *gin.Context) ([]dish.Dish, error) {
	restaurantId := c.Param("restaurantId")
	if restaurantId == "" {
		return []dish.Dish{}, ErrInvalidRestaurantId
	}

	dishes, err := dc.DishService.GetDishesByRestaurantId(c, restaurantId)
	if err != nil {
		return []dish.Dish{}, err
	}
	return dishes, nil
}

func (dc DishControllerImplementation) AddDishesToRestaurant(c *gin.Context, requestBody AddDishesToRestaurantRequestBody) error {
	restaurantId, dishes, err := dc.Marshallers.MarshalAddDishesToRestaurantRequestBody(requestBody)
	if err != nil {
		return err
	}

	err = dc.DishService.AddDishesToRestaurant(c, restaurantId, dishes)
	if err != nil {
		return err
	}
	return nil
}
