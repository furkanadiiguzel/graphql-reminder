package database


 var connectionString = "postgres://username:password@hostname:port/dbname?sslmode=require"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return
	}
	defer db.Close()
