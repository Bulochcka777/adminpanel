package route

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupGetCount(router *gin.Engine) {
	router.POST("/count-widget", handleGetLogCount)
}

func handleGetLogCount(c *gin.Context) {
	var json struct {
		Unique_ID  string `json:"userId"`
		Site_Name  string `json:"name"`
		Goal_Name  string `json:"nameGoal"`
		Widgetname string `json:"nameWidget"`
	}

	if err := c.BindJSON(&json); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	var userID, siteID, goalID, widgetID int

	// Проверка наличия пользователя
	err := db.QueryRow("SELECT ID FROM Users WHERE Unique_ID = $1", json.Unique_ID).Scan(&userID)
	if err == sql.ErrNoRows {
		log.Printf("User with Unique_ID %s not found", json.Unique_ID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		log.Printf("Error querying user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
		return
	}

	// Проверка наличия сайта
	err = db.QueryRow("SELECT ID FROM Site WHERE Name = $1", json.Site_Name).Scan(&siteID)
	if err == sql.ErrNoRows {
		log.Printf("Site with Name %s not found", json.Site_Name)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Site not found"})
		return
	} else if err != nil {
		log.Printf("Error querying site: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query site"})
		return
	}

	// Проверка наличия цели
	err = db.QueryRow("SELECT ID FROM Goal WHERE Name = $1", json.Goal_Name).Scan(&goalID)
	if err == sql.ErrNoRows {
		log.Printf("Goal with Name %s not found", json.Goal_Name)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Goal not found"})
		return
	} else if err != nil {
		log.Printf("Error querying goal: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query goal"})
		return
	}

	// Проверка наличия виджета
	err = db.QueryRow("SELECT ID FROM Widgets WHERE Name = $1", json.Widgetname).Scan(&widgetID)
	if err == sql.ErrNoRows {
		log.Printf("Widget with Name %s not found", json.Widgetname)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Widget not found"})
		return
	} else if err != nil {
		log.Printf("Error querying widget: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query widget"})
		return
	}

	// Подсчет количества логов
	var count int
	query := `
        SELECT COUNT(*)
        FROM Widget_Logs
        WHERE User_ID = $1 AND Site_ID = $2 AND Goal_ID = $3 AND Widget_ID = $4
    `
	err = db.QueryRow(query, userID, siteID, goalID, widgetID).Scan(&count)
	if err != nil {
		log.Printf("Error querying log count: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query log count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}
