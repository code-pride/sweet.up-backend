package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserQueryHandler interface {
	FindById(id primitive.ObjectID) (*User, error)
}

type userQueryHandler struct{}

func NewUserQueryHandler() UserQueryHandler {
	return &userQueryHandler{}
}

func (userCmdHandler *userQueryHandler) FindById(userId primitive.ObjectID) (*User, error) {
	return nil, nil
}
