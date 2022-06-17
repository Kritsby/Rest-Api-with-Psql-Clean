package main

import (
	"rest-api/cmd/app"
	"rest-api/internal/adapters/db/postgreSQL"
	"rest-api/internal/controller/http"
	"rest-api/internal/domain/service"
	"rest-api/pkg/router"
)

//@title API
//@version 1.0
//@description API Server

//@BasePath /
func main() {
	pool := app.App()
	bd := postgreSQL.NewUserStorage(&pool)
	useCase := service.NewTableService(bd)
	c := http.NewFirstHandler(useCase)
	router.NewConnection(c)
}
