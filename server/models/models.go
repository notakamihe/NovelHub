package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey" json:"id"`
	ISBN          string    `gorm:"type:CHARACTER(13);unique_index;NOT NULL" json:"isbn"`
	Title         string    `gorm:"type:VARCHAR(255);NOT NULL" json:"title"`
	Synopsis      string    `gorm:"type:TEXT;NOT NULL" json:"synopsis"`
	PublishDate   string    `gorm:"type:CHARACTER(10);NOT NULL" json:"publishDate"`
	CreatedOn     time.Time `gorm:"type:TIMESTAMP;default:now()" json:"createdOn"`
	PdfUrl        string    `gorm:"type:VARCHAR(255)" json:"pdfUrl"`
	Genre         string    `gorm:"type:VARCHAR(50);NOT NULL" json:"genre"`
	EditionNumber int       `gorm:"type:INTEGER" json:"editionNumber"`
	NumPages      int       `gorm:"type:INTEGER;NOT NULL" json:"numPages"`
	Publisher     string    `gorm:"type:VARCHAR(255)" json:"publisher"`
	CoverUrl      string    `gorm:"type:VARCHAR(255)" json:"coverUrl"`
	Maturity      string    `gorm:"type:VARCHAR(50);NOT NULL" json:"maturity"`
	Reads         int       `gorm:"type:INTEGER;default:0" json:"reads"`
	AuthorID      int
	Author        Author `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Author struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"type:VARCHAR(50);unique_index;NOT NULL" json:"username"`
	Email     string `gorm:"type:VARCHAR(255);unique_index;NOT NULL" json:"email"`
	FirstName string `gorm:"type:VARCHAR(50);NOT NULL" json:"firstName"`
	LastName  string `gorm:"type:VARCHAR(50);NOT NULL" json:"lastName"`
	Password  string `gorm:"type:VARCHAR(255);NOT NULL" json:"password"`
	Title     string `gorm:"type:VARCHAR(255)" json:"title"`
	About     string `gorm:"type:TEXT" json:"about"`
	PfpUrl    string `gorm:"type:VARCHAR(255)" json:"pfpUrl"`
	Dob       string `gorm:"type:CHARACTER(10);NOT NULL" json:"dob"`
}
