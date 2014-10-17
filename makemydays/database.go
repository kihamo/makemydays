package makemydays

import (
	"log"
	"database/sql"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

var dbMap *gorp.DbMap

func GetDatabase() *gorp.DbMap {
	// TODO: options db
	if dbMap == nil {
		db, err := sql.Open("sqlite3", "database.db")
		if err != nil {
			log.Fatalln(err, "sql.Open failed")
		}

		dbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

		dbMap.AddTableWithName(Movie{}, "movies").SetKeys(true, "Id")
		dbMap.AddTableWithName(Song{}, "songs").SetKeys(true, "Id")
		dbMap.AddTableWithName(Word{}, "words").SetKeys(true, "Id")
		dbMap.AddTableWithName(Book{}, "books").SetKeys(true, "Id")
		dbMap.AddTableWithName(Task{}, "tasks").SetKeys(true, "Id")
		dbMap.AddTableWithName(Food{}, "foods").SetKeys(true, "Id")

		if err = dbMap.CreateTablesIfNotExists(); err != nil {
			log.Fatalln("Create tables failed", err)
		}
	}

	return dbMap
}
