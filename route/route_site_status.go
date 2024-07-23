package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Site struct {
	Name      string `json:"name"`
	Working   bool   `json:"working"`
	Debugging bool   `json:"debugging"`
}

// setupRoutes настраивает маршруты HTTP
func setupStatusSite(router *gin.Engine) {
	router.POST("/site_status", handlePostSite)
}

// handlePostSite обрабатывает POST-запросы для получения информации о сайте
func handlePostSite(c *gin.Context) {
	var json struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		log.Printf("Получен некорректный JSON: %v", err)
		return
	}

	siteName := json.Name
	log.Printf("Получен запрос для сайта: %s", siteName)
	if siteName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствует имя сайта"})
		log.Println("Отсутствует имя сайта")
		return
	}

	// Вызываем getSiteFromDB с siteName как параметром
	site, err := getSiteFromDB(siteName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Сайт не найден"})
		log.Printf("Сайт не найден: %s", siteName)
		return
	}

	c.JSON(http.StatusOK, site)
	log.Printf("Отправлен ответ: %+v", site)
}

// getSiteFromDB извлекает данные о сайте из базы данных
func getSiteFromDB(siteName string) (*Site, error) {
	var site Site

	// При необходимости экранируем или обрамляем siteName для SQL-запроса
	safeSiteName := "'" + siteName + "'" // Возможно, вам нужно подстроить это под требования вашей базы данных

	err := db.QueryRow(`SELECT name, working, debugging FROM site WHERE name = `+safeSiteName).Scan(&site.Name, &site.Working, &site.Debugging)
	if err != nil {
		return nil, err
	}
	return &site, nil
}
