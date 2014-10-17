package makemydays

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatalln(err, "sql.Open failed")
	}

	db.AutoMigrate(
		&Movie{},
		&Song{},
		&Word{},
		&Book{},
		&Task{},
		&Food{},
	)

	return &db
}
