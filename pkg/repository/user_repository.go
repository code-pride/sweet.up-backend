package repository

import (
	"context"

	"github.com/code-pride/sweet.up/pkg/core/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("User")
	return &UserRepository{
		collection: collection,
	}
}

func (repo *UserRepository) FindByMail(email string) (user.User, error) {
	var returnedUser user.User
	err := repo.collection.FindOne(context.TODO(), user.User{Email: email}).Decode(&returnedUser)

	return returnedUser, err
}

func (repo *UserRepository) CreateUser(user user.User) {
	repo.collection.InsertOne(context.TODO(), user)
}
