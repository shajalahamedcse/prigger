package handler

import (
    "net/http"
    "github.com/shajalahamedcse/prigger/db"
    "github.com/labstack/echo/v4"
)

type TaskRequest struct {
    Description        string `json:"description"`
    BaremetalPrivateIP string `json:"baremetal_private_ip"`
    BaremetalID        string `json:"baremetal_id"`
}

func CreateTask(c echo.Context) error {
    var taskReq TaskRequest
    if err := c.Bind(&taskReq); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    err := db.InsertTask(taskReq.Description, taskReq.BaremetalPrivateIP, taskReq.BaremetalID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]string{"status": "Task created successfully"})
}
