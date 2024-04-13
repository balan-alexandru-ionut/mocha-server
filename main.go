package main

import (
	"mocha-server/database"
	"mocha-server/server"
)

func main() {
	mongoDatabase := &database.Mongo{}
	mongoDatabase.Connect()

	server.Start()
}
