package main

import (
	"context"
	"log"
	"net/http"

	"github.com/allergeye/surveyor-service/internal/database"
	surveyor_dish "github.com/allergeye/surveyor-service/pkg/surveyor/dish"
	surveyor_restaurant "github.com/allergeye/surveyor-service/pkg/surveyor/restaurant"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery())

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	marshallers := surveyor_restaurant.NewMarshallers()

	dishRepository := database.NewDishRepository(client)
	restaurantRepository := database.NewRestaurantRepository(client)

	dishService := surveyor_dish.NewDishService(sugar, restaurantRepository, dishRepository)
	dishController := surveyor_dish.NewDishContoller(sugar, dishService)
	dishRouter := surveyor_dish.NewDishRouter(sugar, dishController)

	restaurantService := surveyor_restaurant.NewRestaurantService(sugar, restaurantRepository, dishRepository)
	restaurantController := surveyor_restaurant.NewRestaurantController(sugar, restaurantService, marshallers)
	restaurantRouter := surveyor_restaurant.NewRestaurantRouter(sugar, restaurantController)

	restaurant := r.Group("/restaurant")
	{
		restaurant.GET("/", restaurantRouter.GetAllRestaurants)
		restaurant.POST("/", restaurantRouter.AddRestaurant)
	}
	dish := r.Group("/dish")
	{
		dish.GET("/:restaurantId", dishRouter.GetDishesByRestaurantId)
	}

	r.GET("/ping", func(c *gin.Context) {
		sugar.Infof("mainRouter.ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
