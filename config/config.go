package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// SetConfig trigger and set the configuration of the Go project
func SetConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err.Error()))
	}
}
