package database

import (
	"mocha-server/util"
	"mocha-server/util/log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	logger = log.New("database")
	config = util.Configuration()
)

type Database interface {
	Connect()
}

type Mongo struct {
	Client *mongo.Client
}

func (database *Mongo) Connect() {
	// 1. Read username and password from config file
	username := config.Yaml.GetString("database.username")
	password := config.Yaml.GetString("database.password")

	// 2. Build host uri
	host := config.Yaml.GetString("database.host")
	port := config.Yaml.GetString("database.port")

	hostUri := host + ":" + port

	// 3. Build credential options
	credentialOptions := options.Credential{
		Username: username,
		Password: password,
	}

	// 4. Build connection options
	connectionOptions := options.Client().
		SetHosts([]string{hostUri}).
		SetAppName("Mocha").
		SetAuth(credentialOptions)

	// 5. Connect
	context, contextCancel := util.NewBackgroundContext(10 * time.Second)
	defer contextCancel()

	client, err := mongo.Connect(context, connectionOptions)
	if err != nil {
		logger.Panic("Cannot connect to mongoDB", err)
	}

	database.Client = client
}
