package main

import (
	"flag"
	"fmt"

	"github.com/code-pride/sweet.up/pkg/core"
	"github.com/code-pride/sweet.up/pkg/http"
	"github.com/code-pride/sweet.up/pkg/mongorepo"
	"github.com/code-pride/sweet.up/pkg/rest"
	"github.com/code-pride/sweet.up/pkg/util"
	"github.com/gorilla/mux"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello, world.")
	//repository.NewMongoRepository("mongodb://localhost:27017")
	//http.InitHttp()

	confPath := flag.String("conf-path", ".", "yml config file path")

	flag.Parse()

	initApp(initConfiguration(confPath))
}

func initConfiguration(confPath *string) util.Configuration {
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(*confPath) // optionally look for config in the working directory
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil { // Read config and handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var conf util.Configuration

	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("Unable to decode into struct, %v", err))
	}

	return conf
}

func initApp(conf util.Configuration) {
	log := util.ConfigureLogger(conf.Logger)

	defer log.Sync()

	mongoRep := mongorepo.Init(conf.Mongo, log)

	core := core.Init(mongoRep.UserRepo, log)

	ctrl := rest.NewUserQueryCommandController(core.UserCommandHandler, core.UserQueryHandler, log)

	sm := mux.NewRouter()
	ctrl.AttachContoller(sm)
	http.InitHttp(sm, log)
}
