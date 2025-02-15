package route

import (
	"carbon/go-reservation/api/controller"
	"carbon/go-reservation/bootstrap"
	"carbon/go-reservation/domain"
	"carbon/go-reservation/repository"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
    ur := repository.NewUserRepository(db, domain.CollectionUser)
    sc := controller.SignupController{
        Signup
    }
}
