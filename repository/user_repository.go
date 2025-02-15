package repository

import (
	"carbon/go-reservation/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) userRepository {
	return userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
    collection := ur.database.Collection(ur.collection)

    opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
    cursor, err := collection.Find(c, bson.D{}, opts)

    if err != nil {
        return nil, err
    }

    var users []domain.User

    err = cursor.All(c, &users)
    if users == nil {
        return []domain.User{}, err
    }

    return users, err
}
