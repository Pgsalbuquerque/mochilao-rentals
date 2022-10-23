package config

import (
	"sync"

	"github.com/spf13/viper"
)

var currentConfig *Config

var config *viper.Viper = viper.New()

var configOnce *sync.Once = &sync.Once{}

type Config struct {
	MongoConnectionString string
	DBName                string
}

const (
	MongoConnectionString = "MONGO_CONNECTION_STRING"
	DBName                = "DB_NAME"
)

func Get() *Config {
	configOnce.Do(func() {
		config.BindEnv(MongoConnectionString)
		config.BindEnv(DBName)

		currentConfig = &Config{
			MongoConnectionString: config.GetString(MongoConnectionString),
			DBName:                config.GetString(DBName),
		}

	})
	return currentConfig
}
