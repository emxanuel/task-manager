package tasks_repository

import (
	"context"
	task_model "task-manager/internal/models"

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
	cursor, err := r.collection.Find(ctx, nil)
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
