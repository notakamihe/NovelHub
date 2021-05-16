package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/notakamihe/novelhub/server/database"
	"github.com/notakamihe/novelhub/server/models"
	"github.com/notakamihe/novelhub/server/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	dsn := "host=localhost user=postgres password=Akamihe2004! dbname=books port=5432 sslmode=disable"
	database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the database.")
		log.Fatal(err)
	}

	database.DB.AutoMigrate(&models.Book{})
	database.DB.AutoMigrate(&models.Author{})
	database.DB.AutoMigrate(&models.Like{})

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "x-access-token"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r := routes.NewRouter()

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(headers, methods, origins)(r)))
}
