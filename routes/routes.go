package routes

import (
	"adminBDvidj/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouterMain(router *gin.Engine, db *sql.DB) {
	router.GET("/", handlers.ServeTemplate)
	router.POST("/goal", handlers.AddGoal(db))
	router.POST("/update-goal", handlers.UpdateGoal(db))
	router.POST("/delete-goal", handlers.DeleteGoal(db))
	router.POST("/delete-empty-users", handlers.DeleteEmptyRows(db))
	router.POST("/add-site", handlers.AddSite(db))
	router.POST("/delete-site", handlers.DeleteSite(db))
	router.GET("/get-sites", handlers.GetSites(db))
	router.GET("/get-widgets", handlers.GetWidgets(db))
	router.POST("/update-fields", handlers.UpdateFieldsHandler(db))
	router.GET("/fetch_data", handlers.FetchDataHandler(db))
	router.POST("/update-script-param", handlers.UpdateScriptParam(db))
	router.POST("/process", handlers.ProcessHandler(db))
	router.GET("/filter-logs", handlers.FilterLogs(db))
	router.GET("/filter-analytics", handlers.FilterLogsA(db))
}
