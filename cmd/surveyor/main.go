package main

import (
	"context"
	"log"
	"net/http"

	"github.com/allergeye/surveyor-service/internal/database"
	"github.com/allergeye/surveyor-service/internal/lib"
	surveyor "github.com/allergeye/surveyor-service/pkg/surveyor/restaurant"
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

	helpers := lib.NewHelpers()

	dishRepository := database.NewDishRepository(client)

	restaurantRepository := database.NewRestaurantRepository(client)
	restaurantService := surveyor.NewRestaurantService(sugar, restaurantRepository, dishRepository)
	restaurantController := surveyor.NewRestaurantController(sugar, restaurantService, helpers)
	restaurantRouter := surveyor.NewRestaurantRouter(sugar, restaurantController)

	restaurant := r.Group("/restaurant")
	{
		restaurant.GET("/", restaurantRouter.GetAllRestaurants)
		restaurant.POST("/", restaurantRouter.AddRestaurant)
	}

	r.GET("/ping", func(c *gin.Context) {
		sugar.Infof("mainRouter.ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
