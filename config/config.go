package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("config/")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file, ", err)
		return
	}
}
