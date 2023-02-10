package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Name           string             `json:"name"`
	Password       string             `json:"password"`
	Role           string             `json:"role"`
	Organization   string             `json:"organization"`
	Email          string             `json:"email"`
	ProfilePicture string             `json:"picture"`
}
