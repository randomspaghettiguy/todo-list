package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoItemController struct {
	ToDoItemUseCase domain.ToDoItemUseCase
}

func (tc *ToDoItemController) Create(c *gin.Context) {
	var task domain.ToDoItem

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	task.ID = primitive.NewObjectID()
	task.Task = strings.TrimSpace(task.Task)

	if task.Task == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "title cannot be blank"})
		return
	}

	task.Status = false // set to Default

	err = tc.ToDoItemUseCase.Create(c, &task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Item created successfully!",
	})
}

func (tc *ToDoItemController) Fetch(c *gin.Context) {
	tasks, err := tc.ToDoItemUseCase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *ToDoItemController) GetByID(c *gin.Context) {
	taskId := c.Query("x-task-id")

	task, err := tc.ToDoItemUseCase.GetByID(c, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *ToDoItemController) Delete(c *gin.Context) {
	taskId := c.Query("x-task-id")

	err := tc.ToDoItemUseCase.Delete(c, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Deleted item successfully!",
	})
}

func (tc *ToDoItemController) Edit(c *gin.Context) {
	rq := domain.EditTaskRequset{Status: false}
	taskId := c.Query("x-task-id")

	err := c.ShouldBind(&rq)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.ToDoItemUseCase.Edit(c, taskId, rq.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Editted item successfully!",
	})
}
