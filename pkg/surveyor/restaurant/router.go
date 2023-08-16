package surveyor_restaurant

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ErrRestaurantAlreadyExists = errors.New("a restaurant with that name already exists")
	ErrInvalidRestaurantId     = errors.New("invalid restaurant id")
	ErrRestaurantNotFound      = errors.New("a restaurant with that id could not be found")
)

type RestaurantRouter interface {
	GetAllRestaurants(c *gin.Context)
	AddRestaurant(c *gin.Context)
	GetDishesForRestaurant(c *gin.Context)
}

type RestaurantRouterImplementation struct {
	logger               *zap.SugaredLogger
	restaurantController RestaurantController
}

func NewRestaurantRouter(logger *zap.SugaredLogger, restaurantController RestaurantController) RestaurantRouter {
	return RestaurantRouterImplementation{
		logger:               logger,
		restaurantController: restaurantController,
	}
}

func (r RestaurantRouterImplementation) GetAllRestaurants(c *gin.Context) {
	r.logger.Infow("restaurantRouter.GetAllRestaurants")
	restaurants, err := r.restaurantController.GetAllRestaurants(c)
	if err != nil {
		r.logger.Errorw("restaurantRouter.GetAllRestaurants", "error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"restaurants": restaurants,
	})
}

func (r RestaurantRouterImplementation) AddRestaurant(c *gin.Context) {
	r.logger.Infow("restaurantRouter.AddRestaurants")

	var requestBody AddRestaurantRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		r.logger.Errorw("restaurantRouter.AddRestaurants", "error", err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := r.restaurantController.AddRestaurant(c, requestBody)
	if err != nil {
		r.logger.Errorw("restaurantRouter.AddRestaurants", "error", err)
		if errors.Is(err, ErrRestaurantAlreadyExists) {
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

func (r RestaurantRouterImplementation) GetDishesForRestaurant(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	if restaurantId == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	dishes, err := r.restaurantController.GetDishesForRestaurant(c, restaurantId)
	if err != nil {
		if errors.Is(err, ErrInvalidRestaurantId) {
			c.Status(http.StatusBadRequest)
			return
		} else if errors.Is(err, ErrRestaurantNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"dishes": dishes,
	})
}
