package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddGoal(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("goal_name")
		about := c.PostForm("goal_about")
		_, err := db.Exec(`INSERT INTO Goal (Name, About) VALUES ($1, $2)`, name, about)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusFound, "/")
	}
}

func UpdateGoal(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("goal_id")
		name := c.PostForm("goal_name")
		about := c.PostForm("goal_about")
		_, err := db.Exec(`UPDATE Goal SET Name = $1, About = $2 WHERE ID = $3`, name, about, id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusFound, "/")
	}
}

func DeleteGoal(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("goal_id")
		_, err := db.Exec(`DELETE FROM Goal WHERE ID = $1`, id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusFound, "/")
	}
}
