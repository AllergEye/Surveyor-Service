package surveyor_dish

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ErrInvalidAllergen            = errors.New("that is not a valid allergen")
	ErrUserCannotGuessProbability = errors.New("you cannot guess a probability that an allergen is present")
)

type DishRouter interface {
	GetDishesByRestaurantId(c *gin.Context)
	AddDishesToRestaurant(c *gin.Context)
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
	restaurantId := c.Param("restaurantId")
	if restaurantId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	dishes, err := dr.DishController.GetDishesByRestaurantId(c, restaurantId)
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

func (dr DishRouterImplementation) AddDishesToRestaurant(c *gin.Context) {
	var requestBody AddDishesToRestaurantRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		dr.Logger.Errorw("restaurantRouter.AddRestaurants", "error", err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := dr.DishController.AddDishesToRestaurant(c, requestBody)

	if err != nil {
		if errors.Is(err, ErrInvalidAllergen) || errors.Is(err, ErrUserCannotGuessProbability) {
			c.JSON(http.StatusPreconditionFailed, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
