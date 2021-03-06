package config

import (
	"fmt"
	"github.com/ryanbradynd05/go-tmdb"
	"github.com/spf13/viper"
	"strings"
)

var TmdbConfig tmdb.Config

func ReadConfig() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigName("pohls") // name of config file (without extension)
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("$GOPATH/src/github.com/ASVBPREAUBV/pohls/configs")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("phl")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	TmdbConfig = tmdb.Config{
		APIKey:   viper.GetString("tmdbtoken"),
		Proxies:  nil,
		UseProxy: false,
	}

}
