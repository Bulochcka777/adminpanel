package route

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupLogRoutes(router *gin.Engine) {
	router.POST("/add-log", handleAddLog)
}

func handleAddLog(c *gin.Context) {
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
	dvalueStr := convertToString(json.Dvalue)

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

	var userID, siteID, goalID, widgetID, widgetGoalID int

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

	// Проверка наличия цели виджета, с условием на цели без параметра dvalue
	if json.Goal_Name == "closeForm" || json.Goal_Name == "showingForm" || json.Goal_Name == "sendForm" {
		err = tx.QueryRow("SELECT ID FROM Widgets_goal WHERE ID_goal = $1 AND ID_site = $2 AND ID_widgets = $3", goalID, siteID, widgetID).Scan(&widgetGoalID)
	} else {
		err = tx.QueryRow("SELECT ID FROM Widgets_goal WHERE ID_goal = $1 AND Dvalue = $2 AND ID_site = $3 AND ID_widgets = $4", goalID, dvalueStr, siteID, widgetID).Scan(&widgetGoalID)
	}

	if err == sql.ErrNoRows {
		log.Printf("Widget goal not found for Goal ID %d, Site ID %d, Widget ID %d, Dvalue %s", goalID, siteID, widgetID, dvalueStr)
		if json.Goal_Name == "closeForm" || json.Goal_Name == "showingForm" || json.Goal_Name == "sendForm" {
			log.Printf("Note: Goal Name %s does not require a Dvalue", json.Goal_Name)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Widget goal not found"})
		return
	} else if err != nil {
		log.Printf("Error querying widget goal: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query widget goal"})
		return
	}

	// Вставка записи в таблицу Logs
	query := `
        INSERT INTO Logs (User_ID, Visit, ID_widgets_goal, Time)
        VALUES ($1, $2, $3, $4)
    `
	_, err = tx.Exec(query, userID, json.Visit, widgetGoalID, json.Time)
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
func convertToString(value interface{}) string {
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
