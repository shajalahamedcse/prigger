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
    db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName))
    if err != nil {
        panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }
}

func InsertTask(description, baremetalPrivateIP, baremetalID string) error {
    query := `
        INSERT INTO tasks (description, state, baremetal_private_ip, baremetal_id)
        VALUES ($1, 'creating', $2, $3)
    `
    _, err := db.Exec(query, description, baremetalPrivateIP, baremetalID)
    return err
}
