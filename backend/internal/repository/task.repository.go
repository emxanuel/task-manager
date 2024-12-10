package tasks_repository

import (
	"context"
	task_model "task-manager/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewTaskRepository(client *mongo.Client) *Repository {
	db := client.Database("task-manager")
	return &Repository{db.Collection("tasks")}
}

func (r *Repository) FindAll(ctx context.Context) ([]task_model.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var tasks []task_model.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) Create(ctx context.Context, task task_model.Task) (*task_model.Task, error) {
	date := time.Now().Format(time.RFC3339)
	newTask := task_model.Task{
		ID:          primitive.NewObjectID(),
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   date,
	}

	_, err := r.collection.InsertOne(ctx, newTask)

	if err != nil {
		return nil, err
	}

	return &newTask, nil
}
