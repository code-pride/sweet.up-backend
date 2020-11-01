package user

import (
	"github.com/code-pride/sweet.up/pkg/core/apperror"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type UserCommandHandler interface {
	CreateUser(user User) (*primitive.ObjectID, error)

	UpdateUserDetails(userId primitive.ObjectID, userDetails UserDetails) error

	AcceptPair(userId primitive.ObjectID, pairId primitive.ObjectID) error
}

type userCommandHandler struct {
	userRepository UserCommandQueryRepository
	log            *zap.SugaredLogger
}

func NewUserCommandHandler(ur UserCommandQueryRepository, log *zap.SugaredLogger) UserCommandHandler {
	return &userCommandHandler{
		userRepository: ur,
		log:            log,
	}
}

func (uch *userCommandHandler) CreateUser(user User) (*primitive.ObjectID, error) {
	id, err := uch.userRepository.CreateUser(user)
	if err != nil {
		uch.log.Error("Unable to create user: %s", err)
		return nil, err
	}

	return id, nil
}

func (uch *userCommandHandler) UpdateUserDetails(userId primitive.ObjectID, userDetails UserDetails) error {
	usr, err := uch.getUser(userId)
	if err != nil {
		uch.log.Error("Unable to update user: %s", err)
		return err
	}

	usr.UserDetails = userDetails

	return uch.userRepository.UpdateUser(*usr)
}

func (uch *userCommandHandler) AcceptPair(userId primitive.ObjectID, pairId primitive.ObjectID) error {
	usr, err := uch.getUser(userId)
	if err != nil {
		uch.log.Error("Unable to accept pair: %s", err)
		return err
	}

	pair, err := uch.getUser(userId)
	if err != nil {
		uch.log.Error("Unable to accept pair: %s", err)
		return err
	}

	if usr.Pair.ID != primitive.NilObjectID {
		return apperror.NewApplicationError(nil, "User already has a pair")
	}

	if pair.Pair.ID != primitive.NilObjectID {
		return apperror.NewApplicationError(nil, "User pair already has a pair")
	}

	usr.Pair.ID = pairId
	pair.Pair.ID = userId

	uch.userRepository.UpdateUser(*usr)
	uch.userRepository.UpdateUser(*pair)

	return nil
}

func (uch *userCommandHandler) getUser(userId primitive.ObjectID) (*User, error) {
	usr, err := uch.userRepository.FindById(userId)
	if err != nil {
		_, ok := err.(*apperror.EntityNotFoundError)
		if !ok {
			return nil, err
		}
		return nil, nil
	}
	return usr, nil
}
