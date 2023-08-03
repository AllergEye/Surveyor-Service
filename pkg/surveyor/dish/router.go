package surveyor_dish

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DishRouter interface {
	GetDishesByRestaurantId(c *gin.Context)
}

type DishRouterImplementation struct {
	Logger         *zap.SugaredLogger
	DishController DishController
}

func NewDishRouter(logger *zap.SugaredLogger, dishController DishController) DishRouter {
	return DishRouterImplementation{
		Logger:         logger,
		DishController: dishController,
	}
}

func (dr DishRouterImplementation) GetDishesByRestaurantId(c *gin.Context) {
	dishes, err := dr.DishController.GetDishesByRestaurantId(c)
	if err != nil {
		if errors.Is(err, ErrInvalidRestaurantId) {
			c.Status(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"dishes": dishes,
	})
}
