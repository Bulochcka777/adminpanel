package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"adminBDvidj/models"

	"github.com/gin-gonic/gin"
)

func UpdateFieldsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updates []models.SiteUpdate
		if err := c.BindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}
		for _, update := range updates {
			log.Printf("Processing update: %+v", update)
			var query string
			switch update.Field {
			case "working":
				query = `UPDATE Site SET Working = $1 WHERE ID = $2`
			case "debugging":
				query = `UPDATE Site SET Debugging = $1 WHERE ID = $2`
			default:
				log.Printf("Invalid field: %v", update.Field)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid field"})
				return
			}
			_, err := db.Exec(query, update.IsChecked, update.SiteID)
			if err != nil {
				log.Printf("Database update error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database update failed"})
				return
			}
		}
		log.Printf("Updates processed successfully: %+v", updates)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
