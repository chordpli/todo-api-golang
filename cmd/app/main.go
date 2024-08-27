package main

import (
	"log"
	"net/http"
	"todo-api-golang/internal/routes"
	"todo-api-golang/util"

	_ "todo-api-golang/docs" // Swagger docs 패키지 임포트
)

// @title Todo API
// @version 1.0
// @description This is a sample server for managing todos.
// @host localhost:8000
// @BasePath /

// @securityDefinitions.basic BasicAuth
func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server := &http.Server{
		Addr:    config.PORT,
		Handler: routes.Router(),
	}

	log.Fatal(server.ListenAndServe())

}
