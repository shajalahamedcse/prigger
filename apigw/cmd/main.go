package main


import (
    "github.com/shajalahamedcse/prigger/apigw/db"
    handler "github.com/shajalahamedcse/prigger/apigw/handlers"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // Initialize the Echo instance
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Initialize the database
    db.InitDB()

    // Middleware to check for secret key
    e.Use(checkSecretKey)

    // Register routes
    e.POST("/tasks", handler.CreateTask)

    // Start the server
    go func() {
        if err := e.Start(":8080"); err != nil {
            e.Logger.Info("Shutting down the server")
        }
    }()

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down task service...")

    if err := e.Shutdown(nil); err != nil {
        log.Fatal(err)
    }
}

// Middleware function to check for secret key
func checkSecretKey(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        secretKey := c.Request().Header.Get("X-Secret-Key")
        if secretKey != "mysecretkey" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
        }
        return next(c)
    }
}
