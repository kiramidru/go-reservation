package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type userRepository struct {
	database mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository {
		database: db,
		collection: collection,
	}
}
