package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/furkanadiiguzel/graphql-reminder/graph"
	"github.com/furkanadiiguzel/graphql-reminder/graph/generated"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// Load environment variables from the ".env" file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Get the database connection URL from environment variables.
	dbURL := os.Getenv("DB_URL")

	// Connect to the database.
	Database := graph.Connect(dbURL)

	// Create the GraphQL server with the connected database.
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: Database}}))

	// ...
}
