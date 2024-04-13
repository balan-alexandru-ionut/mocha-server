package server

import (
	"mocha-server/util"
	"mocha-server/util/log"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	logger = log.New("server")
	config = util.Configuration()
)

func Start() {
	// 1. Build uri from which to serve
	host := config.Yaml.GetString("server.host")
	port := config.Yaml.GetInt("server.port")

	serverUri := host + ":" + strconv.Itoa(port)

	// 2. Start server
	err := gin.Default().Run(serverUri)
	if err != nil {
		logger.Panic("Could not start server", err)
	}
}
