package app

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"rest-api/pkg/client/postrgeSQL"
)

type DsnConfig struct {
	Name     string
	Password string
	Host     string
	Port     string
	Database string
}

func App() pgxpool.Pool {
	initConfig()

	//go config.StartKafka()
	//
	//fmt.Println("kafka has been started")

	pool := postrgeSQL.InitPostgres()

	return pool
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("cant't find config file")
		}
		panic("can't read config file")
	}

}
