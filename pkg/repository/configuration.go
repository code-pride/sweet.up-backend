package repository

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient(url string) (*MongoClient, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Successfully connected to the ", url)

	return &MongoClient{
		client: client,
	}, nil
}
