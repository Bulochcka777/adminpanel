package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupSiteAliasRoutes(router *gin.Engine) {
	router.POST("/add-site-alias", handleAddSiteAlias)
}

func handleAddSiteAlias(c *gin.Context) {
	var json struct {
		Site struct {
			Name      string `json:"name"`
			Status    string `json:"status"`
			CreatedAt string `json:"created_at"`
			CreatedBy string `json:"created_by"`
			UpdatedAt string `json:"updated_at"`
			UpdatedBy string `json:"updated_by"`
		} `json:"site"`
		Alias struct {
			Domain    string `json:"domain"`
			Subdomain string `json:"subdomain"`
			Tld       string `json:"tld"`
			Flag      bool   `json:"flag"`
		} `json:"alias"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Вставка записи в таблицу Site
	siteQuery := `
        INSERT INTO Site (Name, Status, Created_At, Created_By, Updated_At, Updated_By)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING ID
    `
	var siteID int
	err := db.QueryRow(siteQuery, json.Site.Name, json.Site.Status, json.Site.CreatedAt, json.Site.CreatedBy, json.Site.UpdatedAt, json.Site.UpdatedBy).Scan(&siteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Вставка записи в таблицу Alias
	aliasQuery := `
        INSERT INTO Alias (Site_ID, Domain, Subdomain, Tld, Flag)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err = db.Exec(aliasQuery, siteID, json.Alias.Domain, json.Alias.Subdomain, json.Alias.Tld, json.Alias.Flag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Site and Alias added successfully"})
}
