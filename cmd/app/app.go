package app

import (
	"fmt"
	"github.com/spf13/viper"
	"rest-api/internal/controller/http"
	"rest-api/internal/domain/service"
	"rest-api/internal/domain/usecase/repo"
	"rest-api/pkg/client/postgreSQL"
	"rest-api/pkg/router"
)

func init() {
	initConfig()
}

func Run() error {

	pool, err := postgreSQL.InitPostgres()
	if err != nil {
		return err
	}
	bd := repo.NewUserStorage(pool)
	useCase := service.NewTableService(bd)
	handler := http.NewFirstHandler(useCase)
	handler.Register(router.Router())
	router.NewConnection()

	return err
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
