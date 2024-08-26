package database

import (
	"context"
	"log"
	"todo-api-golang/ent"
	"todo-api-golang/util"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *ent.Client {
	config, err := util.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	dsn := config.RDB

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
