package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func DB_INIT() {
	
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "root")
	dbName := getEnv("DB_NAME", "finn_tracker")

	
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)


	var err error
	maxAttempts := 10
	retryDelay := 5 * time.Second

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err == nil {
			fmt.Println("Successfully connected to the database")
			break
		}

		log.Printf("Failed to connect to the database (Attempt %d/%d): %v", attempts, maxAttempts, err)
		time.Sleep(retryDelay)
	}

	if err != nil {
		log.Fatalf("All connection attempts failed: %v", err)
	}

	if DB == nil {
		log.Fatal("DB is nil after initialization")
	}
}

func Migration(models ...interface{}) {
	for _, model := range models {
		if model == nil {
			log.Fatalf("Migration received a nil model")
		}

		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("Migration failed for model %T: %v", model, err)
		} else {
			fmt.Printf("Migration succeeded for model %T\n", model)
		}
	}
}


func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}