package user

import (
	"errors"

	"github.com/code-pride/sweet.up/pkg/core/apperror"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type UserQueryHandler interface {
	FindById(id primitive.ObjectID) (*User, error)
}

type userQueryHandler struct {
	userRepository UserCommandQueryRepository
	log            *zap.SugaredLogger
}

func NewUserQueryHandler(ur UserCommandQueryRepository, log *zap.SugaredLogger) UserQueryHandler {
	return &userQueryHandler{}
}

func (uqh *userQueryHandler) FindById(userId primitive.ObjectID) (*User, error) {
	usr, err := uqh.userRepository.FindById(userId)
	if err != nil {
		if err != nil {
			if !errors.Is(err, apperror.ErrEntityNotFound) {
				uqh.log.Error(err)
			}
			return nil, err
		}
	}

	return usr, nil
}
