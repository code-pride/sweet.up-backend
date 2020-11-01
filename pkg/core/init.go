package core

import (
	"github.com/code-pride/sweet.up/pkg/core/user"
	"go.uber.org/zap"
)

type Core struct {
	UserCommandHandler user.UserCommandHandler
	UserQueryHandler   user.UserQueryHandler
}

func Init(usrRep user.UserCommandQueryRepository, log *zap.SugaredLogger) *Core {
	return &Core{
		UserCommandHandler: user.NewUserCommandHandler(usrRep, log),
		UserQueryHandler:   user.NewUserQueryHandler(),
	}
}
