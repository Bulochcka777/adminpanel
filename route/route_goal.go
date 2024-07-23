package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupGoalRoutes(router *gin.Engine) {
	router.POST("/add-goal", handleAddGoal)
}

func handleAddGoal(c *gin.Context) {
	var json struct {
		Unique_ID   string `json:"unique_id"`
		Referrer    string `json:"referrer"`
		Device_type string `json:"device_type"`
		Browser     string `json:"browser"`
		Os          string `json:"os"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
        INSERT INTO Users (Unique_ID, Referrer, Device_type, Browser, Os)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := db.Exec(query, json.Unique_ID, json.Referrer, json.Device_type, json.Browser, json.Os)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User added successfully"})
}
