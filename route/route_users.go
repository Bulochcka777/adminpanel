package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupUserRoutes(router *gin.Engine) {
	router.POST("/add-user", handleAddUser)
}

func handleAddUser(c *gin.Context) {
	var json struct {
		Unique_ID   string `json:"unique_id"`
		Reg_flag    string `json:"reg_flag"`
		Reg_id      string `json:"user_reg"`
		Referrer    string `json:"referrer"`
		Device_type string `json:"device_type"`
		Browser     string `json:"browser"`
		Os          string `json:"os"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Добавление отладочной информации для проверки полей
	log.Printf("Parsed JSON: %+v", json)

	query := `
        INSERT INTO Users (Unique_ID, Reg_flag, Reg_id, Referrer, Device_type, Browser, Os)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := db.Exec(query, json.Unique_ID, json.Reg_flag, json.Reg_id, json.Referrer, json.Device_type, json.Browser, json.Os)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User added successfully"})
}
