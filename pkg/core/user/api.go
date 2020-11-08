package user

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Email       string             `bson:"email"`
	UserDetails UserDetails
	Pair        Pair
}

type UserDetails struct {
	Name     string `bson:"name"`
	LastName string `bson:"lastName"`
}

type Pair struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Accepted bool
}

type UpdateUserDetailsCommand struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserDetails UserDetails
}

type UserCommandQueryRepository interface {
	FindByMail(email string) (*User, error)
	FindById(id primitive.ObjectID) (*User, error)
	CreateUser(user User) (*primitive.ObjectID, error)
	UpdateUser(user User) error
	UpdateUsers(users []User) error
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *UserDetails) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (p *Pair) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
