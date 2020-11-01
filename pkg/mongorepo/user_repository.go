package mongorepo

import (
	"context"

	"github.com/code-pride/sweet.up/pkg/core/apperror"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/code-pride/sweet.up/pkg/core/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) user.UserCommandQueryRepository {
	collection := db.Collection("User")
	return &MongoUserRepository{
		collection: collection,
	}
}

func (repo *MongoUserRepository) FindById(id primitive.ObjectID) (*user.User, error) {
	var returnedUser *user.User
	err := repo.collection.FindOne(context.TODO(), user.User{ID: id}).Decode(returnedUser)
	if err != nil {
		return nil, apperror.NewApplicationError(err, err.Error())
	}

	return returnedUser, err
}

func (repo *MongoUserRepository) FindByMail(email string) (*user.User, error) {
	var returnedUser *user.User
	err := repo.collection.FindOne(context.TODO(), user.User{Email: email}).Decode(returnedUser)
	if err != nil {
		return nil, apperror.NewApplicationError(err, err.Error())
	}

	return returnedUser, err
}

func (repo *MongoUserRepository) CreateUser(user user.User) (*primitive.ObjectID, error) {
	result, err := repo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, apperror.NewApplicationError(err, err.Error())
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &oid, nil
	}

	return nil, apperror.NewApplicationError(nil, err.Error())
}

func (repo *MongoUserRepository) UpdateUser(user user.User) error {
	_, err := repo.CreateUser(user)
	return err
}
