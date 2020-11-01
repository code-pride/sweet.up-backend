package util

import (
	"github.com/code-pride/sweet.up/pkg/core/coreutil"
	"github.com/code-pride/sweet.up/pkg/http"
	"github.com/code-pride/sweet.up/pkg/mongorepo"
)

type Configuration struct {
	HTTP   http.HttpConfig
	Core   coreutil.CoreConfiguration
	Mongo  mongorepo.MongoConfiguration
	Logger LoggerConfiguration
}
