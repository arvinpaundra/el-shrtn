package config

import (
	"log"

	"github.com/spf13/viper"
)

var c Config

func LoadEnv(filename, ext, path string) {
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed read env: %s", err.Error())
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("failed unmarshal env: %s", err.Error())
	}
}

func GetAppPort() string {
	return c.AppPort
}

func GetAppEnv() string {
	return c.AppMode
}

func GetBaseUrlApp() string {
	return c.BaseUrlApp
}

func GetMongoURI() string {
	return c.Mongo.URI
}

func GetMongoDBName() string {
	return c.Mongo.DBName
}

func GetMongoUserColl() string {
	return c.Mongo.UserColl
}

func GetMongoLinkColl() string {
	return c.Mongo.LinkColl
}
