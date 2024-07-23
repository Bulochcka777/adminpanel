package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		site := c.PostForm("site")
		widget := c.PostForm("widget")
		goal := c.PostForm("goal")
		dvalue := c.PostForm("dvalue")

		query := "INSERT INTO Widgets_goal (ID_goal, Dvalue, ID_site, ID_widgets) VALUES ($1, $2, $3, $4)"
		fmt.Println("SQL-запрос для вставки данных:")
		fmt.Println(query)
		fmt.Printf("С параметрами: goal=%s, dvalue=%s, site=%s, widget=%s\n", goal, dvalue, site, widget)
		result, err := db.Exec(query, goal, dvalue, site, widget)
		if err != nil {
			fmt.Println("Ошибка при выполнении SQL-запроса:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		lastInsertID, err := result.LastInsertId()
		if err != nil {
			fmt.Println("Ошибка получения ID последней вставки:", err)
		} else {
			fmt.Println("ID последней вставки данных:", lastInsertID)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			fmt.Println("Ошибка получения количества затронутых строк:", err)
		} else {
			fmt.Println("Количество затронутых строк:", rowsAffected)
		}
	}
}
