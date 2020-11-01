package coreutil

import "go.mongodb.org/mongo-driver/bson/primitive"

type CoreConfiguration struct {
	SecurityEnabled bool
}

type EntityCreated struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}
