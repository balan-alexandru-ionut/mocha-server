package main

import (
	"fmt"
	"github.com/spf13/viper"
	"mocha-server/config"
	"mocha-server/server"
)

func main() {
	config.ReadEnvVariables()
	fmt.Println(viper.GetString("test"))

	server.Start()
}
