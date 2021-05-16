package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/notakamihe/novelhub/server/database"
	"github.com/notakamihe/novelhub/server/models"
	"github.com/notakamihe/novelhub/server/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Unscoped().Delete(&models.Author{}, params["id"]); result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Could not delete author w/ id of %s.", params["id"]), 400)
		return
	}

	json.NewEncoder(w).Encode("Author successfully deleted.")
}

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author

	database.DB.Preload(clause.Associations).Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func getAuthorById(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Preload(clause.Associations).First(&author, params["id"]); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Author w/ id of %s does not exist.", params["id"]), 400)
			return
		}

		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(author)
}

func getAuthorByToken(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	headerToken := r.Header.Get("x-access-token")

	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		http.Error(w, err.Error(), 401)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return
	}

	id, ok := claims["user_id"]

	if !ok {
		http.Error(w, "Could not obtain author id.", 401)
		return
	}

	var result *gorm.DB

	if result = database.DB.Preload(clause.Associations).First(&author, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Couldn't find author w/ id of %d.", id), 401)
			return
		}

		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(author)
}

func getAuthorByUsername(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Preload(clause.Associations).Where("username = ?", params["username"]).First(&author); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Author w/ username of %s does not exist.", params["username"]), 400)
			return
		}

		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(author)
}

func getAuthorFans(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Preload(clause.Associations).First(&author, params["id"]); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Author w/ id of %s does not exist.", params["id"]), 400)
			return
		}

		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(author.Fans)
}

func login(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	json.NewDecoder(r.Body).Decode(&author)

	if author.Username == "" {
		http.Error(w, "Must provide a username.", 400)
		return
	}

	password := author.Password

	var result *gorm.DB

	if result = database.DB.Where("username = ?", author.Username).First(&author); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("No user found w/ username of %s.", author.Username), 400)
			return
		}

		log.Fatal(result.Error)
	}

	err := bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(password))

	if err != nil {
		http.Error(w, "Password is incorrect.", 400)
		return
	}

	token, err := utils.CreateToken(author.ID)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(token)
}

func register(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	json.NewDecoder(r.Body).Decode(&author)

	if author.Username == "" {
		http.Error(w, "Must provide a username.", 400)
		return
	} else {
		for _, letter := range author.Username {
			if !unicode.IsLetter(letter) && !unicode.IsDigit(letter) && letter != '_' {
				http.Error(w, "Characters in username must be alphanumeric.", 400)
				return
			}
		}
	}

	if author.Email == "" {
		http.Error(w, "Must provide an email.", 400)
		return
	} else {
		matched, err := regexp.Match(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)`, []byte(author.Email))

		if err != nil || !matched {
			http.Error(w, "Email must be in valid format.", 400)
			return
		}
	}

	if author.FirstName == "" {
		http.Error(w, "Must provide a first name.", 400)
		return
	}

	if author.LastName == "" {
		http.Error(w, "Must provide a last name.", 400)
		return
	}

	if author.Password == "" {
		http.Error(w, "Must provide a password.", 400)
		return
	} else if len(author.Password) < 8 {
		http.Error(w, "Password must be at least eight characters long.", 400)
		return
	}

	if author.Dob == "" {
		http.Error(w, "Must provide a date of birth", 400)
		return
	} else {
		matched, err := regexp.Match(`^[12]\d{3}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])`, []byte(author.Dob))

		if err != nil || !matched {
			http.Error(w, "Invalid date of birth format (Must be YYYY-MM-DD).", 400)
			return
		}

		t, err := time.Parse("2006-01-02", author.Dob)

		if err != nil {
			log.Fatal(err)
		}

		if time.Now().Year()-t.Year() < 13 {
			http.Error(w, "Must be thirteen years or older.", 400)
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(author.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	author.Password = string(hashedPassword)

	var result *gorm.DB

	if result = database.DB.Create(&author); result.Error != nil {
		if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				if strings.Contains(pgError.Message, "idx_authors_username") {
					http.Error(w, "Username must be available.", 400)
					return
				} else if strings.Contains(pgError.Message, "idx_authors_email") {
					http.Error(w, "Email must be available.", 400)
					return
				}
			}
		}
	}

	token, err := utils.CreateToken(author.ID)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(token)
}

func toggleAuthorFan(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	var fan models.Author
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Preload(clause.Associations).First(&author, params["id"]); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Author w/ id of %s does not exist.", params["id"]), 400)
			return
		}

		log.Fatal(result.Error)
	}

	if result = database.DB.Preload(clause.Associations).First(&fan, params["fanId"]); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Fan w/ id of %s does not exist.", params["fanId"]), 400)
			return
		}

		log.Fatal(result.Error)
	}

	for _, f := range author.Fans {
		if f.ID == fan.ID {
			database.DB.Model(&author).Association("Fans").Delete([]*models.Author{&fan})
			database.DB.Preload(clause.Associations).First(&author, params["id"])

			json.NewEncoder(w).Encode(author)
			return
		}
	}

	database.DB.Model(&author).Association("Fans").Append([]*models.Author{&fan})
	database.DB.Preload(clause.Associations).First(&author, params["id"])

	json.NewEncoder(w).Encode(author)
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&author)

	if author.Username == "" {
		http.Error(w, "Must provide a username.", 400)
		return
	} else {
		for _, letter := range author.Username {
			if !unicode.IsLetter(letter) && !unicode.IsDigit(letter) && letter != '_' {
				http.Error(w, "Characters in username must be alphanumeric.", 400)
				return
			}
		}
	}

	if author.Email == "" {
		http.Error(w, "Must provide an email.", 400)
		return
	} else {
		matched, err := regexp.Match(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)`, []byte(author.Email))

		if err != nil || !matched {
			http.Error(w, "Email must be in valid format.", 400)
			return
		}
	}

	if author.FirstName == "" {
		http.Error(w, "Must provide a first name.", 400)
		return
	}

	if author.LastName == "" {
		http.Error(w, "Must provide a last name.", 400)
		return
	}

	if author.Dob == "" {
		http.Error(w, "Must provide a date of birth", 400)
		return
	} else {
		matched, err := regexp.Match(`^[12]\d{3}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])`, []byte(author.Dob))

		if err != nil || !matched {
			http.Error(w, "Invalid date of birth format (Must be YYYY-MM-DD).", 400)
			return
		}

		t, err := time.Parse("2006-01-02", author.Dob)

		if err != nil {
			log.Fatal(err)
		}

		if time.Now().Year()-t.Year() < 13 {
			http.Error(w, "Must be thirteen years or older.", 400)
			return
		}
	}

	result := database.DB.Model(&author).Where("id = ?", params["id"]).Updates(map[string]interface{}{
		"username":   author.Username,
		"email":      author.Email,
		"first_name": author.FirstName,
		"last_name":  author.LastName,
		"dob":        author.Dob,
		"title":      author.Title,
		"about":      author.About,
	})

	if result.Error != nil {
		if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				if strings.Contains(pgError.Message, "idx_authors_username") {
					http.Error(w, "Username must be available.", 400)
					return
				} else if strings.Contains(pgError.Message, "idx_authors_email") {
					http.Error(w, "Email must be available.", 400)
					return
				}
			}
		}
	}

	if result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Could not update author w/ id of %s.", params["id"]), 400)
		return
	}

	database.DB.Preload(clause.Associations).First(&author, params["id"])

	json.NewEncoder(w).Encode(author)
}

func updateAuthorPfp(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	pfpUrl := ""
	params := mux.Vars(r)

	r.ParseMultipartForm(16 << 20)

	file, handler, err := r.FormFile("pfp")

	if err != nil {
		if file != nil {
			http.Error(w, "Error retrieving file.", 400)
			return
		}
	}

	if file != nil {
		defer file.Close()
	}

	if handler != nil {
		if !strings.Contains(handler.Header.Get("Content-Type"), "image") {
			http.Error(w, "Invalid media type. Must be image.", 400)
			return
		}

		pfpUrl = filepath.Join("images", time.Now().Format("20060102150405")+handler.Filename)
	}

	result := database.DB.Model(&author).Where("id = ?", params["id"]).Update("pfp_url", pfpUrl)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Could not update pfp url of author w/ id of %s.", params["id"]), 400)
		return
	}

	if handler != nil && file != nil {
		path := filepath.Join("uploads", pfpUrl)

		newFile, err := os.Create(path)

		if err != nil {
			log.Fatal(err)
		}

		defer newFile.Close()

		bytes, err := ioutil.ReadAll(file)

		if err != nil {
			log.Fatal(err)
		}

		newFile.Write(bytes)
	}

	database.DB.Preload(clause.Associations).First(&author, params["id"])

	json.NewEncoder(w).Encode(author)
}
