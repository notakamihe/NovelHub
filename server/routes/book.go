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
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/notakamihe/novelhub/server/database"
	"github.com/notakamihe/novelhub/server/models"
	"github.com/notakamihe/novelhub/server/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var genres = []string{
	"fiction", "narrative", "science fiction", "non-fiction", "mystery", "historcial fiction", "biography",
	"horror", "poem", "humor", "short", "autobiography", "romance", "crime", "fairy tale/folklore", "moral",
	"fantasy", "comic", "essay", "dystopian", "drama", "satire", "memoir", "suspense", "other",
}

var maturities = []string{"childrens", "teen", "young adult", "adult", "any"}

func addReadToBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	result := database.DB.Model(&book).Where("id = ?", params["id"]).Update("reads", gorm.Expr("reads + 1"))

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Could not update book w/ id of %s.", params["id"]), 400)
		return
	}

	database.DB.Preload(clause.Associations).First(&book, params["id"])

	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	if book.ISBN == "" {
		http.Error(w, "Book ISBN must be specified", 400)
		return
	} else if len(book.ISBN) != 13 {
		http.Error(w, "Valid ISBNs must be 13 characters long.", 400)
		return
	}

	if book.Title == "" {
		http.Error(w, "Must provide a title.", 400)
		return
	}

	if book.Synopsis == "" {
		http.Error(w, "Must provide a synopsis.", 400)
		return
	} else {
		words := strings.Split(book.Synopsis, " ")

		if len(words) < 15 {
			http.Error(w, "Book synopsis must be at least 15 words long.", 400)
			return
		}
	}

	if book.PublishDate == "" {
		book.PublishDate = time.Now().Format("2006-01-02")
	} else {
		matched, err := regexp.Match(`^[12]\d{3}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])`, []byte(book.PublishDate))

		if err != nil || !matched {
			http.Error(w, "Invalid format for publish date (Must be YYYY-MM-DD).", 400)
			return
		}
	}

	if book.Genre == "" {
		http.Error(w, "Must provide a genre.", 400)
		return
	} else if !utils.SliceContainsIgnoreCase(genres, book.Genre) {
		http.Error(w, "Genre must be one of these categories: "+strings.Join(genres, ", "), 400)
		return
	}

	if book.NumPages <= 0 {
		http.Error(w, "Must provide a page quantity greater than 0.", 400)
		return
	}

	if book.Maturity == "" {
		book.Maturity = "any"
	} else if !utils.SliceContainsIgnoreCase(maturities, book.Maturity) {
		http.Error(w, "Maturity must be one of these categories: "+strings.Join(maturities, ", "), 400)
		return
	}

	var result *gorm.DB

	if result = database.DB.Create(&book); result.Error != nil {
		if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				if strings.Contains(pgError.Message, "idx_books_isbn") {
					http.Error(w, "ISBN not available.", 400)
					return
				}

				break
			case "23503":
				if strings.Contains(pgError.Message, "fk_books_author") {
					http.Error(w, fmt.Sprintf("Author ID of %d is not valid.", book.AuthorID), 400)
					return
				}

				break
			}
		}

		log.Fatal(result.Error)
	}

	database.DB.Preload(clause.Associations).First(&book, book.ID)

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Unscoped().Delete(&models.Book{}, params["id"]); result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Could not delete book w/ id of %s.", params["id"]), 400)
		return
	}

	json.NewEncoder(w).Encode("Book successfully deleted.")
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	var conditions []string

	params := r.URL.Query()

	search := params.Get("search")
	genre := params.Get("genre")
	fromDate := params.Get("fromDate")
	toDate := params.Get("toDate")
	comparison := params.Get("comparison")
	numPages := params.Get("numPages")
	maturity := params.Get("maturity")

	if genre != "" && utils.SliceContainsIgnoreCase(genres, genre) {
		conditions = append(conditions, fmt.Sprintf("genre = '%s'", strings.ToLower(genre)))
	}

	if fromDate != "" {
		matched, err := regexp.Match(`^[12]\d{3}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])`, []byte(fromDate))

		if err == nil && matched {
			conditions = append(conditions, fmt.Sprintf("publish_date >= '%s'", fromDate))
		}
	}

	if toDate != "" {
		matched, err := regexp.Match(`^[12]\d{3}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])`, []byte(toDate))

		if err == nil && matched {
			conditions = append(conditions, fmt.Sprintf("publish_date <= '%s'", toDate))
		}
	}

	if numPages != "" {
		if comparison != "" {
			switch comparison {
			case "equals":
				conditions = append(conditions, fmt.Sprintf("num_pages = %s", numPages))
				break
			case "lt":
				conditions = append(conditions, fmt.Sprintf("num_pages < %s", numPages))
				break
			case "lte":
				conditions = append(conditions, fmt.Sprintf("num_pages <= %s", numPages))
				break
			case "gt":
				conditions = append(conditions, fmt.Sprintf("num_pages > %s", numPages))
				break
			case "gte":
				conditions = append(conditions, fmt.Sprintf("num_pages >= %s", numPages))
				break
			}
		}
	}

	if maturity != "" && utils.SliceContainsIgnoreCase(maturities, maturity) {
		conditions = append(conditions, fmt.Sprintf("maturity = '%s'", strings.ToLower(maturity)))
	}

	var conditionsStr string

	if len(conditions) >= 1 {
		conditionsStr = strings.Join(conditions, " and ")
	} else {
		conditionsStr = "1 = 1"
	}

	SEARCH_QUERY := fmt.Sprintf(
		`
	books.isbn like '%%%[1]v%%' or
	lower(books.title) like lower('%%%[1]v%%') or 
	lower(books.publisher) like lower('%%%[1]v%%') or
	lower(authors.username) like lower('%%%[1]v%%') or
	lower(concat(authors.first_name, ' ', authors.last_name)) like lower('%%%[1]v%%')
	`, search)

	QUERY := fmt.Sprintf(
		"select books.* from books inner join authors on books.author_id = authors.id where %s and (%s)",
		conditionsStr, SEARCH_QUERY)

	database.DB.Preload(clause.Associations).Raw(QUERY).Find(&books)
	json.NewEncoder(w).Encode(books)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Preload(clause.Associations).First(&book, params["id"]); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Book w/ id of %s does not exist.", params["id"]), 400)
			return
		}

		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(book)
}

func getBooksByAuthorId(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Preload(clause.Associations).Where("author_id", params["id"]).Find(&books); result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(books)
}

func getLikedBooksByAuthorId(w http.ResponseWriter, r *http.Request) {
	var likes []models.Like
	params := mux.Vars(r)

	var result *gorm.DB

	if result = database.DB.Preload(clause.Associations).Preload("Book."+clause.Associations).Where("author_id", params["id"]).Find(&likes); result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(likes)
}

func toggleBookLike(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var like models.Like
	params := mux.Vars(r)

	bookId, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	authorId, err := strconv.Atoi(params["authorId"])

	if err != nil {
		log.Fatal(err)
	}

	var result *gorm.DB

	if result = database.DB.Where("book_id = ? and author_id = ?", bookId, authorId).First(&like); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newLike := models.Like{BookID: uint(bookId), AuthorID: uint(authorId)}

			if result = database.DB.Create(&newLike); result.Error != nil {
				if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
					switch pgError.Code {
					case "23503":
						if strings.Contains(pgError.Message, "fk_likes_author") {
							http.Error(w, fmt.Sprintf("Author ID of %d is not valid.", authorId), 400)
							return
						}

						if strings.Contains(pgError.Message, "fk_books_likes") {
							http.Error(w, fmt.Sprintf("Book ID of %d is not valid.", bookId), 400)
							return
						}

						break
					}
				}

				log.Fatal(result.Error)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		if result = database.DB.Delete(&like); result.Error != nil {
			log.Fatal(result.Error)
		}
	}

	database.DB.Preload(clause.Associations).First(&book, params["id"])

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&book)

	if book.ISBN == "" {
		http.Error(w, "Book ISBN must be specified", 400)
		return
	} else if len(book.ISBN) != 13 {
		http.Error(w, "Valid ISBNs must be 13 characters long.", 400)
		return
	}

	if book.Title == "" {
		http.Error(w, "Must provide a title.", 400)
		return
	}

	if book.Synopsis == "" {
		http.Error(w, "Must provide a synopsis.", 400)
		return
	} else {
		words := strings.Split(book.Synopsis, " ")

		if len(words) < 15 {
			http.Error(w, "Book synopsis must be at least 15 words long.", 400)
			return
		}
	}

	if book.PublishDate == "" {
		book.PublishDate = time.Now().Format("2006-01-02")
	} else {
		matched, err := regexp.Match(`^[12]\d{3}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])`, []byte(book.PublishDate))

		if err != nil || !matched {
			http.Error(w, "Invalid format for publish date (Must be YYYY-MM-DD).", 400)
			return
		}
	}

	if book.Genre == "" {
		http.Error(w, "Must provide a genre.", 400)
		return
	} else if !utils.SliceContainsIgnoreCase(genres, book.Genre) {
		http.Error(w, "Genre must be one of these categories: "+strings.Join(genres, ", "), 400)
		return
	}

	if book.NumPages <= 0 {
		http.Error(w, "Must provide a page quantity greater than 0.", 400)
		return
	}

	if book.Maturity == "" {
		book.Maturity = "any"
	} else if !utils.SliceContainsIgnoreCase(maturities, book.Maturity) {
		http.Error(w, "Maturity must be one of these categories: "+strings.Join(maturities, ", "), 400)
		return
	}

	result := database.DB.Model(&book).Where("id = ?", params["id"]).Updates(map[string]interface{}{
		"isbn":           book.ISBN,
		"title":          book.Title,
		"synopsis":       book.Synopsis,
		"publish_date":   book.PublishDate,
		"genre":          book.Genre,
		"edition_number": book.EditionNumber,
		"num_pages":      book.NumPages,
		"publisher":      book.Publisher,
		"maturity":       book.Maturity,
	})

	if result.Error != nil {
		if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				if strings.Contains(pgError.Message, "idx_books_isbn") {
					http.Error(w, "ISBN not available.", 400)
					return
				}

				break
			case "23503":
				if strings.Contains(pgError.Message, "fk_books_author") {
					http.Error(w, fmt.Sprintf("Author ID of %d is not valid.", book.AuthorID), 400)
					return
				}

				break
			}
		}

		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Could not update book w/ id of %s.", params["id"]), 400)
		return
	}

	database.DB.Preload(clause.Associations).First(&book, params["id"])

	json.NewEncoder(w).Encode(book)
}

func updateBookCover(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	coverUrl := ""
	params := mux.Vars(r)

	r.ParseMultipartForm(16 << 20)

	file, handler, err := r.FormFile("cover")

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

		coverUrl = filepath.Join("images", time.Now().Format("20060102150405")+handler.Filename)
	}

	result := database.DB.Model(&book).Where("id = ?", params["id"]).Update("cover_url", coverUrl)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Could not update cover url of book w/ id of %s.", params["id"]), 400)
		return
	}

	if handler != nil && file != nil {
		path := filepath.Join("uploads", coverUrl)

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

	database.DB.Preload(clause.Associations).First(&book, params["id"])

	json.NewEncoder(w).Encode(book)
}

func updateBookPdf(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	pdfUrl := ""
	params := mux.Vars(r)

	r.ParseMultipartForm(16 << 20)

	file, handler, err := r.FormFile("pdf")

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
		if handler.Header.Get("Content-Type") != "application/pdf" {
			http.Error(w, "Invalid file type. Must be PDF.", 400)
			return
		}

		pdfUrl = filepath.Join("pdfs", time.Now().Format("20060102150405")+handler.Filename)
	}

	result := database.DB.Model(&book).Where("id = ?", params["id"]).Update("pdf_url", pdfUrl)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Could not update pdf url of book w/ id of %s.", params["id"]), 400)
		return
	}

	if handler != nil && file != nil {
		path := filepath.Join("uploads", pdfUrl)

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

	database.DB.Preload(clause.Associations).First(&book, params["id"])

	json.NewEncoder(w).Encode(book)
}
