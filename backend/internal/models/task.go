package task_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	CreatedAt   string             `bson:"createdAt,omitempty"`
}

type CreateTaskBody struct {
	Title      string `json:"title" binding:"required"`
	Decription string `json:"description"`
}
