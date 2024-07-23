package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWidgets(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		widgetsa, err := getWidgetsFromDB(db)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get widgets: %v", err)
			return
		}
		c.JSON(http.StatusOK, widgetsa)
	}
}
