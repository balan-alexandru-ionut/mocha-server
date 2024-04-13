package util

import (
	"mocha-server/util/log"

	"github.com/spf13/viper"
)

type configuration struct {
	Environment *viper.Viper
	Yaml        *viper.Viper
}

var (
	configurationInstance *configuration = nil
	environmentViper                     = viper.New()
	yamlViper                            = viper.New()
	logger                               = log.New("configuration")
)

func Configuration() *configuration {
	if configurationInstance != nil {
		return configurationInstance
	} else {
		environmentViper.SetEnvPrefix("mocha")
		environmentViper.AutomaticEnv()

		yamlViper.SetConfigName("config")
		yamlViper.SetConfigType("yaml")
		yamlViper.AddConfigPath("$HOME/.config/mocha")
		yamlViper.AddConfigPath("./")
		err := yamlViper.ReadInConfig()
		if err != nil {
			logger.Panic(err)
		}

		configurationInstance = &configuration{
			Environment: environmentViper,
			Yaml:        yamlViper,
		}

		return configurationInstance
	}
}
