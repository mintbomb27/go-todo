package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	ID          primitive.ObjectID  `json:"_id" bson:"_id"`
	TimeCreated primitive.Timestamp `json:"timeCreated"`
	CreatedBy   string              `json:"createdBy"`
	Title       string              `json:"name"`
	Description string              `json:"status"`
	DateTime    primitive.Timestamp `json:"dateTime"`
}
