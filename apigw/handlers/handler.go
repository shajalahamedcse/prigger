package handlers

import (
    "net/http"
    "github.com/shajalahamedcse/prigger/apigw/db"
    "github.com/labstack/echo/v4"
)

type TaskRequest struct {
    UserID             string `json:"user_id"`
    Description        string `json:"description"`
    BaremetalPrivateIP string `json:"baremetal_private_ip"`
    BaremetalID        string `json:"baremetal_id"`
}

// CreateTask handles the creation of a new task.
func CreateTask(c echo.Context) error {
    var taskReq TaskRequest
    
    // Binding the request body to the TaskRequest struct
    if err := c.Bind(&taskReq); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    // Inserting the new task into the database
    id, state, err := db.InsertTask(taskReq.UserID,taskReq.Description, taskReq.BaremetalPrivateIP, taskReq.BaremetalID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    // Returning a successful response with the task details
    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": "Task created successfully",
        "task_id": id,
        "state": state,
    })
}
