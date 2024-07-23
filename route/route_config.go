package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Sites struct {
	ID   int    `json:"-"`
	Name string `json:"name"`
}

type Goal struct {
	ID        int         `json:"-"`
	Name      string      `json:"name"`
	Parameter interface{} `json:"parameter"`
}

type Widgets struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	PathHTML string `json:"html"`
	PathCSS  string `json:"css"`
	PathJS   string `json:"js"`
}

type WidgetsGoal struct {
	ID        int         `json:"-"`
	IDGoal    int         `json:"-"`
	DValue    interface{} `json:"parameter"`
	IDSite    int         `json:"-"`
	IDWidgets int         `json:"-"`
}

type DisplayCondition struct {
	Goals []Goal `json:"goals"`
}

type Resources struct {
	HTML string `json:"html"`
	CSS  string `json:"css"`
	JS   string `json:"js"`
}

type Widget struct {
	Name             string           `json:"name"`
	DisplayCondition DisplayCondition `json:"displayCondition"`
	Resources        Resources        `json:"resources"`
}

type SiteConfig struct {
	Widgets map[string]Widget `json:"widgets"`
}

func setupConfig(router *gin.Engine) {
	router.POST("/config", handleconfig)
}

func handleconfig(c *gin.Context) {
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

	// Вызываем getSiteFromDBconfig с siteName как параметром
	siteConfig, err := getSiteFromDBconfig(siteName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Сайт не найден"})
		log.Printf("Сайт не найден: %s", siteName)
		return
	}

	c.JSON(http.StatusOK, siteConfig)
	log.Printf("Отправлен ответ для сайта %s: %+v", siteName, siteConfig)
}

func getSiteFromDBconfig(siteName string) (SiteConfig, error) {
	var siteConfig SiteConfig
	siteConfig.Widgets = make(map[string]Widget)

	// Добавляем одинарные кавычки вокруг имени сайта
	safeSiteName := "'" + siteName + "'"

	var sites Sites
	err := db.QueryRow(`SELECT ID, Name FROM Site WHERE Name = `+safeSiteName).Scan(&sites.ID, &sites.Name)
	if err != nil {
		return siteConfig, err
	}

	rows, err := db.Query(`
		SELECT wg.Dvalue, g.Name, w.Name, w.path_html, w.path_css, w.path_js
		FROM Widgets_goal wg
		JOIN Goal g ON wg.ID_goal = g.ID
		JOIN Widgets w ON wg.ID_widgets = w.ID
		WHERE wg.ID_site = $1
	`, sites.ID)
	if err != nil {
		return siteConfig, err
	}
	defer rows.Close()

	for rows.Next() {
		var wg WidgetsGoal
		var g Goal
		var w Widgets
		err := rows.Scan(&wg.DValue, &g.Name, &w.Name, &w.PathHTML, &w.PathCSS, &w.PathJS)
		if err != nil {
			return siteConfig, err
		}

		if _, exists := siteConfig.Widgets[w.Name]; !exists {
			siteConfig.Widgets[w.Name] = Widget{
				Name: w.Name,
				DisplayCondition: DisplayCondition{
					Goals: []Goal{},
				},
				Resources: Resources{
					HTML: w.PathHTML,
					CSS:  w.PathCSS,
					JS:   w.PathJS,
				},
			}
		}

		widget := siteConfig.Widgets[w.Name]
		g.Parameter = wg.DValue
		widget.DisplayCondition.Goals = append(widget.DisplayCondition.Goals, g)
		siteConfig.Widgets[w.Name] = widget
	}

	return siteConfig, nil
}
