package main

import (
	"fmt"
	"os"

	dbconnect "github.com/kevinkimutai/metadata/internal/adapter/db/dbConnect"
	"github.com/kevinkimutai/metadata/internal/adapter/server"
	application "github.com/kevinkimutai/metadata/internal/app/core/api"
)

func main() {
	// //Get env var in development
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env files")
	// }

	// Get database connection details from environment variables
	POSTGRES_USERNAME := os.Getenv("POSTGRES_USERNAME")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	PORT := os.Getenv("APPLICATION_PORT")
	DATABASE_PORT := os.Getenv("DATABASE_PORT")

	//Concatinate DB String
	DBURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		POSTGRES_USERNAME,
		POSTGRES_PASSWORD,
		"localhost",
		DATABASE_PORT,
		"moviedb")

	//Connect To DB
	dbAdapter := dbconnect.NewDB(DBURL)

	//Application
	application := application.NewApplication(dbAdapter)

	//Server
	server := server.New(PORT, application)
	server.Run()
}
