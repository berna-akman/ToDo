package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config Configuration

type Configuration struct {
	*Postgres
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SSLMode  string `mapstructure:"sslMode"`
	Port     string `mapstructure:"port"`
}

func InitConfig() (*Postgres, error) {
	// TODO: set env
	viper.SetConfigName("secrets")
	viper.SetConfigType("json")
	viper.AddConfigPath("pg/resource")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("file changed: %s\n", in.Name)
	})
	viper.WatchConfig()

	err = viper.Unmarshal(&config)
	if err != nil {
		return Configuration{}.Postgres, err // TODO: decide for default config if any
	}
	return config.Postgres, err
}
