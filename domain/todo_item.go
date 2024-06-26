package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Collection = "Items"
)

type ToDoItem struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status"`
}

type ToDoItemRepository interface {
	Create(c context.Context, task *ToDoItem) error
	Fetch(c context.Context) ([]ToDoItem, error)
	GetByID(c context.Context, id string) (ToDoItem, error)
}

type ToDoItemUseCase interface {
	Create(c context.Context, task *ToDoItem) error
	Fetch(c context.Context) ([]ToDoItem, error)
	GetByID(c context.Context, id string) (ToDoItem, error)
}
