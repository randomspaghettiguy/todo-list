package repository

import (
	"context"

	"github.com/randomspaghettiguy/go-backend-clean-architecture/domain"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type toDoItemRepository struct {
	database   mongo.Database
	collection string
}

func NewToDoItemRepository(db mongo.Database, collection string) domain.ToDoItemRepository {
	return &toDoItemRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *toDoItemRepository) Create(c context.Context, task *domain.ToDoItem) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, task)

	return err
}

func (tr *toDoItemRepository) Fetch(c context.Context) ([]domain.ToDoItem, error) {
	collection := tr.database.Collection(tr.collection)

	opts := options.Find()
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var items []domain.ToDoItem

	err = cursor.All(c, &items)
	if err != nil {
		return []domain.ToDoItem{}, err
	}

	return items, err
}

func (tr *toDoItemRepository) GetByID(c context.Context, id string) (domain.ToDoItem, error) {
	collection := tr.database.Collection(tr.collection)

	var item domain.ToDoItem

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return item, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&item)
	return item, err
}
