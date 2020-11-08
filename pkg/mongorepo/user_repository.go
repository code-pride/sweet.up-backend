package mongorepo

import (
	"context"
	"errors"

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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperror.ErrEntityNotFound
		}
		return nil, apperror.NewDatabaseError(err)
	}

	return returnedUser, err
}

func (repo *MongoUserRepository) FindByMail(email string) (*user.User, error) {
	var returnedUser *user.User
	err := repo.collection.FindOne(context.TODO(), user.User{Email: email}).Decode(returnedUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperror.ErrEntityNotFound
		}
		return nil, apperror.NewDatabaseError(err)
	}

	return returnedUser, err
}

func (repo *MongoUserRepository) CreateUser(user user.User) (*primitive.ObjectID, error) {
	result, err := repo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, apperror.NewDatabaseError(err)
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &oid, nil
	}

	return nil, apperror.NewInternalServerError("Drvier returned invalid value")
}

func (repo *MongoUserRepository) UpdateUser(user user.User) error {
	_, err := repo.CreateUser(user)
	return err
}

func (repo *MongoUserRepository) UpdateUsers(users []user.User) error {
	u := make([]interface{}, len(users))
	for i, v := range users {
		u[i] = v
	}
	_, err := repo.collection.InsertMany(context.TODO(), u)
	if err != nil {
		return apperror.NewDatabaseError(err)
	}

	return err
}
