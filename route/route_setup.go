package route

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "bdvidj"
)

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
}
func SetupRouter(router *gin.Engine) {
	setupConfig(router)
	setupLogRouteswidget(router)
	setupStatusSite(router)
	setupLogRoutes(router)
	setupUpdateUserRegistration(router)
	setupSiteAliasRoutes(router)
	setupUserRoutes(router)
	setupGoalRoutes(router)
	setupGetCount(router)
}
