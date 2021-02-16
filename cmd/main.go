package main

import (
	"log"

	"github.com/Kuzmrom7/ping-go/config"
	"github.com/Kuzmrom7/ping-go/pkg"
	"github.com/spf13/viper"
)

func main() {
	/* Config parse */
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	var cfg *config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	/* Run screenshot job */
	pkg.RunJob(pkg.Run, cfg)
}
