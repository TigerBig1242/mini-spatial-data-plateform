package database

import (
	"context"
	"fmt"
	"time"

	// "log"
	// "os"

	"github.com/tigerbig/spatial-data-plateform/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	// "gorm.io/gorm/logger"
)

func ConnectDatabase(config *config.Config) (*mongo.Client, error) {
	// func ConnectDatabase(uri string) (*mongo.Client, error) {
	// client, err := mongo.Connect(options.Client().ApplyURI(uri))

	// Connect to database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s uri=%s sslmode=disable",
		config.DB_Host, config.DB_Port, config.DB_User, config.DB_Password,
		config.Uri)
	fmt.Println("DSN Connection string: ", dsn)

	// New logger for detailed SQL logging
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log Level
	// 		Colorful:      true,        // Enable Color
	// 	},
	// )

	client, err := mongo.Connect(options.Client().ApplyURI(config.Uri))
	// client, err := mongo.Connect(options.Client().ApplyURI(uri))
	// fmt.Println("MONGODB_URI MongoDB File", client)
	if err != nil {
		panic(err)
		// return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	errPing := client.Ping(ctx, nil)
	if errPing != nil {
		panic(errPing)
	}

	fmt.Println("Connected to MongoDB")
	fmt.Println("🔥 MIGRATE RUN")

	// defer func() {
	// 	err := client.Disconnect(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()

	return client, nil
	// return client, nil
}
