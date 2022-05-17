package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func Config(key string) string {

	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("./")

	// Alternative of those three above
	viper.SetConfigFile("./config.yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("error on reading config file %s \n", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("invalid type assertion for : %v \n", key)
	}
	return value

}

func IntConfig(key string) int {

	value := Config(key)

	k, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("invalid type assertion for : %v \n", key)
	}
	return k

}
