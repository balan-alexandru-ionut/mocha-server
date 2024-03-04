package config

import "github.com/spf13/viper"

func ReadEnvVariables() {
	viper.SetEnvPrefix("mocha")
	viper.AutomaticEnv()
}
