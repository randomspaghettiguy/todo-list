package usecase

import (
	"context"
	"time"

	"github.com/randomspaghettiguy/go-backend-clean-architecture/domain"
)

type ToDoItemUseCase struct {
	toDoItemRepository domain.ToDoItemRepository
	contextTimeout     time.Duration
}

func NewToDoItemUseCase(toDoItemRepository domain.ToDoItemRepository, timeout time.Duration) domain.ToDoItemUseCase {
	return &ToDoItemUseCase{
		toDoItemRepository: toDoItemRepository,
		contextTimeout:     timeout,
	}
}

func (tu *ToDoItemUseCase) Create(c context.Context, task *domain.ToDoItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.toDoItemRepository.Create(ctx, task)
}

func (tu *ToDoItemUseCase) Fetch(c context.Context) ([]domain.ToDoItem, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.toDoItemRepository.Fetch(ctx)
}

func (tu *ToDoItemUseCase) GetByID(c context.Context, id string) (domain.ToDoItem, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.toDoItemRepository.GetByID(ctx, id)
}
