package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var TmdbToken string

func readConfigFile() {
	viper.SetConfigName("pohls_config") // name of config file (without extension)
	viper.AddConfigPath("config")       // optionally look for config in the working directory

	viper.SetEnvPrefix("phl")
	viper.BindEnv("tmbd_token")

	TmdbToken = viper.GetString("tmbd_token")

	TmdbToken = "döfslgkj"

	fmt.Println(TmdbToken)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
