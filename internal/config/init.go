package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

var TmdbToken string

func ReadConfig() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigName("pohls") // name of config file (without extension)
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("phl")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	TmdbToken = viper.GetString("tmdbtoken")

}
