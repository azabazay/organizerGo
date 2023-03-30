package main

import (
	"log"
	"os"

	"github.com/abai/organizer/storage"
	types "github.com/abai/organizer/types"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = types.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate User")
	}

	err = types.MigrateTimeTableItem(db)
	if err != nil {
		log.Fatal("could not migrate TimeTableItem")
	}

	svc := NewUserService(
		"https://catfact.ninja/fact",
		db,
	)
	svc = NewLoggingService(svc)

	apiServer := NewApiServer(svc)
	log.Fatal(apiServer.Start(":3000"))
}
