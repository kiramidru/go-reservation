package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name string             `bson:"first_name"`
	Last_name  string             `bson:"last_name"`
	Password   string             `bson:"password"`
	Email      string             `bson:"email"`
	Phone      string             `bson:"phone"`
	Token      string             `bson:"token"`
	User_type  string             `bson:"refresh_token"`
	Created_at time.Time          `bson:"created_at"`
	Updated_at time.Time          `bson:"updated_at"`
}

type UserRepository interface {
	create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
