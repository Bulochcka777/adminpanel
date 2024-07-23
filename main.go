package main

import (
	"adminBDvidj/handlers"
	"adminBDvidj/route"
	"adminBDvidj/routes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB
var tmpl = template.Must(template.ParseFiles("template.html"))

func main() {
	gin.SetMode(gin.ReleaseMode)
	var err error
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	f, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))
	route.SetupRouter(router)
	routes.SetupRouterMain(router, db)
	router.Static("/static", "./static")
	router.Static("/widgets", "./widgets")
	handlers.InitTemplateHandler(tmpl, db)
	certFile := "server.crt"
	keyFile := "server.key"
	err = router.RunTLS(":8080", certFile, keyFile)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
