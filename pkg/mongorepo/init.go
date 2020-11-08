package mongorepo

import (
	"context"

	"github.com/code-pride/sweet.up/pkg/core/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoConfiguration struct {
	Host string
	Db   string
}

type MongoRepository struct {
	db     *mongo.Database
	client *mongo.Client

	// Repositories
	UserRepo user.UserCommandQueryRepository
}

func Init(conf MongoConfiguration, log *zap.SugaredLogger) MongoRepository {
	client := createMongoClient(conf.Host, log)
	db := client.Database(conf.Db)

	return MongoRepository{
		db:       db,
		client:   client,
		UserRepo: NewUserRepository(db),
	}
}

func createMongoClient(url string, log *zap.SugaredLogger) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Panicf(err.Error())
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Panicf(err.Error())
	}

	log.Debug("Successfully initialized mongo repository ", url)

	return client
}
