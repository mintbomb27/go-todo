package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"todo-backend/database"
	"todo-backend/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateToDo(c *fiber.Ctx) error {
	var data map[string]string
	var ToDo models.ToDo
	// var userID primitive.ObjectID

	// err := CheckIfAuthorized(c, &userID)
	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		fmt.Println(err.Error())
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ToDo.ID = primitive.NewObjectID()
	//Finding User name
	// var creator models.User
	// usersCollection := database.MongoClient.Database("todoDB").Collection("users")
	// err = usersCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&creator)
	// if err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": err,
	// 	})
	// }

	ToDo.CreatedBy = "anonymous"
	ToDo.Title = data["title"]
	dateUnix, _ := strconv.Atoi(data["dateTime"])
	ToDo.DateTime = primitive.Timestamp{T: uint32(dateUnix)}
	ToDo.Description = data["description"]
	ToDo.TimeCreated = primitive.Timestamp{T: uint32(time.Now().Unix())}

	collection := database.MongoClient.Database("todoDB").Collection("ToDos")

	_, err := collection.InsertOne(ctx, ToDo)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetToDos(c *fiber.Ctx) error {
	collection := database.MongoClient.Database("todoDB").Collection("ToDos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	var ToDos []bson.M
	if err = cursor.All(ctx, &ToDos); err != nil {
		log.Fatal(err)
	}

	return c.JSON(ToDos)
}
