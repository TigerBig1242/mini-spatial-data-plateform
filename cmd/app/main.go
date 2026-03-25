package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tigerbig/spatial-data-plateform/internal/config"
	"github.com/tigerbig/spatial-data-plateform/internal/infrastructure/database"
)

func main() {
	databaseConfig := config.LoadConfig()

	client, err := database.ConnectDatabase(databaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("spatial-data")

	migrateErr := database.Migrate(db)
	if migrateErr != nil {
		panic(migrateErr)
	}
	fmt.Println("🚀 APP START", time.Now().UnixNano())

	// defer func() {
	// 	err := client.Disconnect(db)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()

}
