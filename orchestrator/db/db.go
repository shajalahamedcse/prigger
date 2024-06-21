package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
	"os"
)

var db *sql.DB

func InitDB() {
    var err error
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    
    // Opening the database connection
    db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName))
    if err != nil {
        panic(fmt.Sprintf("Error opening database: %v", err))
    }

    // Verifying the connection
    if err = db.Ping(); err != nil {
        panic(fmt.Sprintf("Error connecting to the database: %v", err))
    }
}

// InsertTask inserts a new task into the database and returns the task ID and state.
func InsertTask(userID, description, baremetalPrivateIP, baremetalID string) (int, string, error) {
    sqlStatement := `
	INSERT INTO tasks ( description, state, baremetal_private_ip, baremetal_id)
    VALUES ($1, 'create', $2, $3)
    RETURNING id, state`
    
    var id int
    var state string

    // Executing the query and scanning the returned id and state
    err := db.QueryRow(sqlStatement, description, baremetalPrivateIP, baremetalID).Scan(&id, &state)
    if err != nil {
        return 0, "", fmt.Errorf("error inserting task: %v", err)
    }

    fmt.Printf("New task ID is %d with initial state %s\n", id, state)
    return id, state, nil
}
