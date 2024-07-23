package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSite(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("site_name")
		working := c.PostForm("site_working") == "true"
		debugging := c.PostForm("site_debugging") == "true"
		domain := c.PostForm("domain")
		var siteID int
		err := db.QueryRow(`INSERT INTO Site (Name, Working, Debugging) VALUES ($1, $2, $3) RETURNING id`, name, working, debugging).Scan(&siteID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		_, err = db.Exec(`INSERT INTO Alias (site_id, domain) VALUES ($1, $2)`, siteID, domain)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func DeleteSite(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		siteID := c.PostForm("site_id")
		_, err := db.Exec(`DELETE FROM Alias WHERE site_id = $1`, siteID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		_, err = db.Exec(`DELETE FROM Site WHERE id = $1`, siteID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func GetSites(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sites, err := getSitesFromDB(db)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get sites: %v", err)
			return
		}
		c.JSON(http.StatusOK, sites)
	}
}
