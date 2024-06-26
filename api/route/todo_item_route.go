package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/api/controller"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/bootstrap"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/domain"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/mongo"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/repository"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/usecase"
)

func NewToDoItemRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewToDoItemRepository(db, domain.Collection)
	tc := &controller.ToDoItemController{
		ToDoItemUseCase: usecase.NewToDoItemUseCase(tr, timeout),
	}
	group.POST("/task", tc.Create)
	group.GET("/tasks", tc.Fetch)
	group.GET("/task/id", tc.GetByID)
}
