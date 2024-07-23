package route

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Для использования PostgreSQL
)

func setupLogRouteswidget(router *gin.Engine) {
	router.POST("/add-log-widget", handleAddLogwidget)
}

func handleAddLogwidget(c *gin.Context) {
	var json struct {
		Unique_ID  string      `json:"unique_id"`
		Reg_flag   interface{} `json:"reg_flag"`
		Reg_ID     string      `json:"reg_id"`
		Visit      string      `json:"visit"`
		Site_Name  string      `json:"site_name"`
		Domain     string      `json:"domain"`
		Goal_Name  string      `json:"goal_name"`
		Dvalue     interface{} `json:"dvalue"`
		Widgetname string      `json:"name_widget"`
		Time       string      `json:"time"`
	}

	if err := c.BindJSON(&json); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	// Преобразование Dvalue в строку
	dvalueStr := convertToStringwidg(json.Dvalue)

	// Добавление отладочной информации для проверки полей
	log.Printf("Parsed JSON: %+v", json)

	// Проверка наличия всех необходимых полей в JSON
	if json.Unique_ID == "" || json.Site_Name == "" || json.Goal_Name == "" || json.Visit == "" || json.Time == "" {
		log.Printf("Unique_ID: %s, Site_Name: %s, Goal_Name: %s, Visit: %s, Time: %s", json.Unique_ID, json.Site_Name, json.Goal_Name, json.Visit, json.Time)
		log.Println("One or more required fields are empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (unique_id, site_name, goal_name, visit, time) must be provided"})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}
	defer tx.Rollback()

	var userID, siteID, goalID, widgetID int

	// Проверка наличия пользователя
	err = tx.QueryRow("SELECT ID FROM Users WHERE Unique_ID = $1", json.Unique_ID).Scan(&userID)
	if err == sql.ErrNoRows {
		log.Printf("User with Unique_ID %s not found", json.Unique_ID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		log.Printf("Error querying user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
		return
	}

	// Проверка наличия сайта
	err = tx.QueryRow("SELECT ID FROM Site WHERE Name = $1", json.Site_Name).Scan(&siteID)
	if err == sql.ErrNoRows {
		log.Printf("Site with Name %s not found", json.Site_Name)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Site not found"})
		return
	} else if err != nil {
		log.Printf("Error querying site: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query site"})
		return
	}

	// Проверка наличия цели
	err = tx.QueryRow("SELECT ID FROM Goal WHERE Name = $1", json.Goal_Name).Scan(&goalID)
	if err == sql.ErrNoRows {
		log.Printf("Goal with Name %s not found", json.Goal_Name)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Goal not found"})
		return
	} else if err != nil {
		log.Printf("Error querying goal: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query goal"})
		return
	}

	// Проверка наличия виджета
	err = tx.QueryRow("SELECT ID FROM Widgets WHERE Name = $1", json.Widgetname).Scan(&widgetID)
	if err == sql.ErrNoRows {
		log.Printf("Widget with Name %s not found", json.Widgetname)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Widget not found"})
		return
	} else if err != nil {
		log.Printf("Error querying widget: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query widget"})
		return
	}

	// Вставка записи в таблицу Widget_Logs
	query := `
        INSERT INTO Widget_Logs (User_ID, Visit, Site_ID, Goal_ID, Widget_ID, Time, UNIQUE_ID, REG_FLAG, REG_ID, DOMAIN, DVALUE)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `
	_, err = tx.Exec(query, userID, json.Visit, siteID, goalID, widgetID, json.Time, json.Unique_ID, json.Reg_flag, json.Reg_ID, json.Domain, dvalueStr)
	if err != nil {
		log.Printf("Error inserting log: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Log added successfully"})
}

// Функция для преобразования значения в строку
func convertToStringwidg(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}
