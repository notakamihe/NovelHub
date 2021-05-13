package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/notakamihe/novelhub/server/database"
	"github.com/notakamihe/novelhub/server/models"
)

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author

	database.DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func register(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	json.NewDecoder(r.Body).Decode(&author)

	result := database.DB.Create(&author)

	if result.Error != nil {
		fmt.Println(result)
		log.Fatal(result.Error)
	}

	fmt.Println(result)
}
