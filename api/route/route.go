package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/bootstrap"
	"github.com/randomspaghettiguy/go-backend-clean-architecture/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRoute := gin.Group("")

	NewToDoItemRoute(env, timeout, db, publicRoute)
}
