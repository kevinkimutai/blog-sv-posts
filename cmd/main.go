package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	dbconnect "github.com/kevinkimutai/metadata/internal/adapter/db/dbConnect"
	"github.com/kevinkimutai/metadata/internal/adapter/server"
	application "github.com/kevinkimutai/metadata/internal/app/core/api"
)

func main() {
	//Get env var
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	PORT := os.Getenv("APPLICATION_PORT")

	// Get database connection details from environment variables
	DBURL := os.Getenv("DATABASE_URL")

	//Connect To DB
	dbAdapter := dbconnect.NewDB(DBURL)

	//Application
	application := application.NewApplication(dbAdapter)

	//Server
	server := server.New(PORT, application)
	server.Run()
}
