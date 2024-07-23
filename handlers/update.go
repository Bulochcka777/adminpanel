package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateScriptParam(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		siteId := c.PostForm("site")
		widgetId := c.PostForm("widget")
		goalId := c.PostForm("goal")
		dvalue := c.PostForm("dvalue")
		if id == "" || siteId == "" || widgetId == "" || goalId == "" {
			c.String(http.StatusBadRequest, "Один или несколько параметров пусты")
			return
		}
		_, err := db.Exec(`UPDATE Widgets_goal SET id_site = $1, id_widgets = $2, id_goal = $3, dvalue = $4 WHERE id = $5`, siteId, widgetId, goalId, dvalue, id)
		if err != nil {
			fmt.Println("Ошибка при обновлении записи:", err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
