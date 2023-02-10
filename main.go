package main

import (
	"context"
	"fmt"
	"os"
	"todo-backend/database"
	"todo-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	database.Connect()

	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Print(port)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	if err := database.MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	if port == "" {
		app.Listen(":5000")
	}
	app.Listen(":" + port)
}
