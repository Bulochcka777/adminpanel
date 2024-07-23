package handlers

import (
	"database/sql"
	"net/http"

	"adminBDvidj/models"

	"github.com/gin-gonic/gin"
)

func FetchDataHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sitesf := []models.Sitef{}
		widgets := []models.Widget{}
		goalsf := []models.Goalf{}
		siteRows, err := db.Query("SELECT id, name FROM Site")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer siteRows.Close()
		for siteRows.Next() {
			var sitef models.Sitef
			if err := siteRows.Scan(&sitef.ID, &sitef.Name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			sitesf = append(sitesf, sitef)
		}
		widgetRows, err := db.Query("SELECT id, name FROM Widgets")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer widgetRows.Close()
		for widgetRows.Next() {
			var widget models.Widget
			if err := widgetRows.Scan(&widget.ID, &widget.Name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			widgets = append(widgets, widget)
		}
		goalRows, err := db.Query("SELECT id, name FROM Goal")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer goalRows.Close()
		for goalRows.Next() {
			var goalf models.Goalf
			if err := goalRows.Scan(&goalf.ID, &goalf.Name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			goalsf = append(goalsf, goalf)
		}
		tableRows, err := db.Query(`
			SELECT t.id, s.name as site, w.name as widget, g.name as goal, t.Dvalue as dvalue
			FROM Widgets_goal t
			JOIN Site s ON t.ID_site = s.id
			JOIN Widgets w ON t.ID_widgets = w.id
			JOIN Goal g ON t.ID_goal = g.id
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer tableRows.Close()
		var tableData []models.TableData
		for tableRows.Next() {
			var row models.TableData
			if err := tableRows.Scan(&row.ID, &row.Site, &row.Widget, &row.Goal, &row.Dvalue); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			tableData = append(tableData, row)
		}
		response := models.ResponseData{
			Sitesf:    sitesf,
			Widgets:   widgets,
			Goalsf:    goalsf,
			TableData: tableData,
		}
		c.JSON(http.StatusOK, response)
	}
}
