package route

import (
	"carbon/go-reservation/bootstrap"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, router *gin.Engine) {
    publicRouter := router.Group("")
    // Public APIs
    NewSignupRouter(env, timeout, db, publicRouter)

    privateRouter := router.Group("")
    log.Print(privateRouter)
}
