package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupUpdateUserRegistration(router *gin.Engine) {
	router.POST("/update-user-reg", handleUpdateUser)
}

func handleUpdateUser(c *gin.Context) {
	var json struct {
		Unique_ID string `json:"unique_id"`
		Reg_flag  string `json:"reg_flag"`
		User_reg  string `json:"user_reg"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
        UPDATE Users
        SET Reg_id = $1, Reg_flag = $2
        WHERE Unique_ID = $3
    `
	_, err := db.Exec(query, json.User_reg, json.Reg_flag, json.Unique_ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User updated successfully"})
}
