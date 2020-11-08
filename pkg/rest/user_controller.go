package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"github.com/code-pride/sweet.up/pkg/core/apperror"
	"github.com/code-pride/sweet.up/pkg/core/user"
	"go.uber.org/zap"
)

type UserQueryCommandController struct {
	userCommandHandler user.UserCommandHandler
	userQueryHandler   user.UserQueryHandler
	log                *zap.SugaredLogger
}

func NewUserQueryCommandController(uch user.UserCommandHandler, uqh user.UserQueryHandler, log *zap.SugaredLogger) *UserQueryCommandController {
	return &UserQueryCommandController{
		log:                log,
		userCommandHandler: uch,
		userQueryHandler:   uqh,
	}
}

func (c *UserQueryCommandController) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ur := user.User{}

	err := ur.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse request", http.StatusBadRequest)
		return
	}
	uc, err := c.userCommandHandler.CreateUser(ur)
	if err != nil {
		http.Error(rw, "Unable to create new User", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(uc)
}

func (c *UserQueryCommandController) UpdateUserDetails(rw http.ResponseWriter, r *http.Request) {
	ud := user.UserDetails{}

	err := ud.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse request", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = c.userCommandHandler.UpdateUserDetails(id, ud)
	if err != nil {
		http.Error(rw, "Unable to update User details", http.StatusInternalServerError)
		return
	}
}

func (c *UserQueryCommandController) AcceptPair(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idUsr, err := primitive.ObjectIDFromHex(vars["idUsr"])
	if err != nil {
		http.Error(rw, "Unable to convert user id", http.StatusBadRequest)
		return
	}

	idPair, err := primitive.ObjectIDFromHex(vars["idPair"])
	if err != nil {
		http.Error(rw, "Unable to convert pair id", http.StatusBadRequest)
		return
	}

	err = c.userCommandHandler.AcceptPair(idUsr, idPair)
	if err != nil {
		badReqErr, ok := err.(*apperror.UserReqError)
		if ok {
			http.Error(rw, badReqErr.Error(), http.StatusBadRequest)
			return
		}

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserQueryCommandController) FindById(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	usr, err := c.userQueryHandler.FindById(id)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			http.Error(rw, err.Error(), http.StatusNotFound)
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(usr)
}

func (c *UserQueryCommandController) AttachContoller(sm *mux.Router) {
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/user", c.CreateUser)

	putRouter := sm.Methods(http.MethodPost).Subrouter()
	putRouter.HandleFunc("/user/{id}", c.UpdateUserDetails)
	putRouter.HandleFunc("/user/{idUsr}/pair/{idPair}", c.AcceptPair)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/user/{is}", c.FindById)
}
