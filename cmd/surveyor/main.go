package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/allergeye/surveyor-service/internal/database"
	surveyor_dish "github.com/allergeye/surveyor-service/pkg/surveyor/dish"
	surveyor_restaurant "github.com/allergeye/surveyor-service/pkg/surveyor/restaurant"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func main() {
	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("could not load .env file")
		}
	}

	r := gin.New()

	r.Use(gin.Recovery())

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()

	connString := fmt.Sprintf("%v://%v:%v@%v/?%v", os.Getenv("DB_CONN_METHOD"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_CONN_OPTIONS"))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	restaurantMarshallers := surveyor_restaurant.NewMarshallers()
	dishMarshallers := surveyor_dish.NewMarshallers()

	dishRepository := database.NewDishRepository(client)
	restaurantRepository := database.NewRestaurantRepository(client)

	dishService := surveyor_dish.NewDishService(sugar, restaurantRepository, dishRepository)
	dishController := surveyor_dish.NewDishContoller(sugar, dishService, dishMarshallers)
	dishRouter := surveyor_dish.NewDishRouter(sugar, dishController)

	restaurantService := surveyor_restaurant.NewRestaurantService(sugar, restaurantRepository, dishRepository)
	restaurantController := surveyor_restaurant.NewRestaurantController(sugar, restaurantService, restaurantMarshallers)
	restaurantRouter := surveyor_restaurant.NewRestaurantRouter(sugar, restaurantController)

	restaurant := r.Group("/restaurant")
	{
		restaurant.GET("/", restaurantRouter.GetAllRestaurants)
		restaurant.GET("/dishes/:restaurantId", restaurantRouter.GetDishesForRestaurant)
		restaurant.POST("/", restaurantRouter.AddRestaurant)
	}
	dish := r.Group("/dish")
	{
		dish.POST("/", dishRouter.AddDishesToRestaurant)
	}

	r.GET("/ping", func(c *gin.Context) {
		sugar.Infof("mainRouter.ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
