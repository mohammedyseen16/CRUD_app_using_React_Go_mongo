package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie  string             `json:"movie,omitempty"`
	Status bool               `json:"status,omitempty"`
}
