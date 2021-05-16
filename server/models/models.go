package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN          string `gorm:"type:CHARACTER(13);uniqueIndex;NOT NULL" json:"isbn"`
	Title         string `gorm:"type:VARCHAR(255);NOT NULL" json:"title"`
	Synopsis      string `gorm:"type:TEXT;NOT NULL" json:"synopsis"`
	PublishDate   string `gorm:"type:CHARACTER(10);NOT NULL" json:"publishDate"`
	PdfUrl        string `gorm:"type:VARCHAR(255)" json:"pdfUrl"`
	Genre         string `gorm:"type:VARCHAR(50);NOT NULL" json:"genre"`
	EditionNumber int    `gorm:"type:INTEGER" json:"editionNumber"`
	NumPages      int    `gorm:"type:INTEGER;NOT NULL" json:"numPages"`
	Publisher     string `gorm:"type:VARCHAR(255)" json:"publisher"`
	CoverUrl      string `gorm:"type:VARCHAR(255)" json:"coverUrl"`
	Maturity      string `gorm:"type:VARCHAR(50);NOT NULL" json:"maturity"`
	Reads         int    `gorm:"type:INTEGER;default:0" json:"reads"`
	AuthorID      int    `json:"authorID"`
	Author        Author `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;preload:true" json:"author"`
	Likes         []Like `gorm:"foreignKey:BookID" json:"likes"`
}

type Author struct {
	gorm.Model
	Username  string    `gorm:"type:VARCHAR(50);uniqueIndex;NOT NULL" json:"username"`
	Email     string    `gorm:"type:VARCHAR(255);uniqueIndex;NOT NULL" json:"email"`
	FirstName string    `gorm:"type:VARCHAR(50);NOT NULL" json:"firstName"`
	LastName  string    `gorm:"type:VARCHAR(50);NOT NULL" json:"lastName"`
	Password  string    `gorm:"type:VARCHAR(255);NOT NULL" json:"password"`
	Title     string    `gorm:"type:VARCHAR(255)" json:"title"`
	About     string    `gorm:"type:TEXT" json:"about"`
	PfpUrl    string    `gorm:"type:VARCHAR(255)" json:"pfpUrl"`
	Dob       string    `gorm:"type:CHARACTER(10);NOT NULL" json:"dob"`
	Books     []Book    `gorm:"foreignKey:AuthorID;" json:"books"`
	Fans      []*Author `gorm:"many2many:user_fans" json:"fans"`
}

type Like struct {
	gorm.Model
	BookID   uint   `json:"bookID"`
	Book     Book   `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"book"`
	AuthorID uint   `json:"authorID"`
	Author   Author `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"author"`
}
