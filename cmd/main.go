package main

import (
	"fmt"
	"rest-api/cmd/app"
)

//@title API
//@version 1.0
//@description API Server

//@BasePath /
func main() {
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
