# GraphQL Reminder

This repository contains a GraphQL server for managing reminder listings. The server is built using [gqlgen](https://github.com/99designs/gqlgen), a popular GraphQL library for Go. The application allows users to create, update, and delete reminder listings.

## Project Structure

### `server.go`

The main entry point of the GraphQL server. It initializes the server, connects to the database, and sets up the GraphQL schema. The server is configured to read environment variables from a `.env` file, and it retrieves the database connection URL (`DB_URL`) from these variables.

### `tools.go`

This file is used for managing external tools with the `//go:build tools` build tag. In this case, it includes the `gqlgen` tool.

### `database/database.go`

The database package contains the logic for connecting to the PostgreSQL database. It reads the database connection URL from the environment variables, establishes a connection, and checks the database status.

### `graph/schema.resolvers.go`

This file contains the GraphQL resolver implementations. Resolvers handle the actual execution of GraphQL queries and mutations. The implemented resolvers include:

- `CreateReminderListing`: Creates a new reminder listing.
- `UpdateReminderListing`: Updates an existing reminder listing.
- `DeleteReminderListing`: Deletes a reminder listing.
- `Reminders`: Retrieves all reminder listings.
- `Reminder`: Retrieves a specific reminder listing by ID.

The resolvers use the database connection established in the `database.go` file.

## How to Run

1. **Clone the Repository:**
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Set Up the Environment:**
   Create a `.env` file with the following content:
   ```env
   DB_URL=<your-database-url>
   ```

3. **Run the Server:**
   ```bash
   go run server.go
   ```

4. **Explore GraphQL API:**
   Open your browser and go to `http://localhost:8080/` to access the GraphQL Playground. Here, you can interact with the GraphQL API and test different queries and mutations.

## Note

Feel free to explore and modify the code based on your requirements. This repository serves as a starting point for building GraphQL APIs in Go using `gqlgen`.
