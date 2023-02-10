package controllers

import (
	"context"
	"time"
	"todo-backend/database"
	"todo-backend/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SecretKey = "abracadabra"

func Hello(c *fiber.Ctx) error {
	return c.SendString("It's up!! Yayy! 🚀")
}

func Register(c *fiber.Ctx) error {
	var user models.User
	// var data map[string]string

	if err := c.BodyParser(&user); err != nil {
		return c.SendString(err.Error())
	}
	user.ID = primitive.NewObjectID()

	// password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	// user.Email = data["email"]
	// user.Name = data["name"]
	// user.Organization = data["organization"]
	// user.ProfilePicture = data["picture"]
	// user.Role = data["role"]
	// user.Password = password

	//TODO: check if user already exists with email
	collection := database.MongoClient.Database("todoDB").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, map[string]string{"email": user.Email}).Decode(&user)
	if err == nil { //if user exists
		return c.JSON(fiber.Map{
			"message": "user exists",
		})
	}
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.SendString(err.Error())
	}

	var user models.User

	collection := database.MongoClient.Database("todoDB").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, map[string]string{"email": data["email"]}).Decode(&user)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "unable to authenticate",
		})
	}
	// fmt.Println(user.Password)
	// fmt.Println(bcrypt.GenerateFromPassword([]byte(data["password"]), 12))
	// if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	c.JSON(fiber.Map{
	// 		"message": "incorrect password",
	// 	})
	// 	return
	// }

	if data["password"] != user.Password { //have to hash and match
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "unable to authenticate",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.ID.Hex(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not create JWT",
		})
	}

	cookie := fiber.Cookie{
		Name:     "ieeejwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func CheckIfAuthorized(c *fiber.Ctx, userID *primitive.ObjectID) error {
	cookie := c.Cookies("ieeejwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	*userID, err = primitive.ObjectIDFromHex(claims.Issuer)
	if err != nil {
		return err
	}
	return nil
}
