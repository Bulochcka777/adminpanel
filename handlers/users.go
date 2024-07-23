package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteEmptyRows(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := deleteEmptyRows(db)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, "Empty rows deleted successfully.")
	}
}
