package main

import (
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/furkanadiiguzel/graphql-reminder/graph"
	"github.com/furkanadiiguzel/graphql-reminder/graph/generated"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	// Load environment variables from the "database.env" file.
	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatal("Error loading database.env file:", err)
	}

	// Get the database connection settings from environment variables.
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct the database connection URL.
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to the database.
	Database := graph.Connect(dbURL)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: Database}}))

}
