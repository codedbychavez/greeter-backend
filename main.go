package main

import (
	"greeter-backend/api/controllers"
	"greeter-backend/api/middleware"
	"greeter-backend/api/routing"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)



func main() {
	// Setup env
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")

	// Error checking for .env file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Error, could not locate .env file", err)
		} else {
			fmt.Println("Error loading .env file", err)
		}
	}

	// Configure cors
	var corsConfig = cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
	}

	// Setup error handling
	fiberConfig := fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	}

	loggerConfig := fiberLogger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}

	app := fiber.New(fiberConfig)
	app.Use(fiberLogger.New(loggerConfig))
	app.Use(fiberRecover.New())
	app.Use(cors.New(corsConfig))

	controllerManager := controllers.NewControllerManager()

	// Setup routing
	baseRouter := app.Group("/api/v1")
	routing.CreatePublicRoutes(baseRouter, controllerManager)

	env, ok := viper.Get("ENVIRONMENT").(string)

	if !ok {
		err := errors.New("ENVIRONMENT not found")
		fmt.Println(err)
	}

	fmt.Println("Starting server via", "environment", env)

	startServer(app)
}

func startServer(app *fiber.App) {
	port, ok := viper.Get("PORT").(string)

	if !ok {
		err := errors.New("PORT not found")
		fmt.Println(err)
	}

	listenErr := app.Listen(":" + port)

	if listenErr != nil {
		fmt.Println("Error during listen", listenErr)
	}
}