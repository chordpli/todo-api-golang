package main

import (
	"context"
	"log"
	"todo-api-golang/ent"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *ent.Client {
	dsn := "root:0000@tcp(localhost:3306)/todo?parseTime=True"
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
