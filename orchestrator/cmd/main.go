package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    _ "github.com/lib/pq"
    "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "myuser"
    password = "mypassword"
    dbname   = "mydb"
)

var (
    psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    go listenToPostgres()

    e.Logger.Fatal(e.Start(":8081"))
}

func listenToPostgres() {
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    listener := pq.NewListener(psqlInfo, 1*time.Second, time.Minute, reportProblem)
    err = listener.Listen("task_channel")
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case notification := <-listener.Notify:
            if notification != nil {
                log.Println("Received data from channel [", notification.Channel, "] :", notification.Extra)
            }
        case <-time.After(90 * time.Second):
            go func() {
                listener.Ping()
            }()
        }
    }
}

func reportProblem(ev pq.ListenerEventType, err error) {
    if err != nil {
        fmt.Println(err.Error())
    }
}
