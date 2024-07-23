package handlers

import (
	"adminBDvidj/models"
	"database/sql"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var tmpl *template.Template
var db *sql.DB

func InitTemplateHandler(template *template.Template, database *sql.DB) {
	tmpl = template
	db = database
}

func ServeTemplate(c *gin.Context) {
	logs, err := getLogs(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get logs: %v", err)
		return
	}
	goals, err := getGoals(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get goals: %v", err)
		return
	}
	users, err := getUsers(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get users: %v", err)
		return
	}
	sites, err := getSitesFromDB(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get sites: %v", err)
		return
	}
	widget_Logs, err := getLogsAnalit(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get logs: %v", err)
		return
	}
	data := struct {
		Logs        []models.Logs
		Goals       []models.Goal
		Users       []models.Users
		Sites       []models.Site
		Widget_Logs []models.Widget_Logs
	}{
		Logs:        logs,
		Goals:       goals,
		Users:       users,
		Sites:       sites,
		Widget_Logs: widget_Logs,
	}
	c.Writer.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to render template: %v", err)
	}
}
