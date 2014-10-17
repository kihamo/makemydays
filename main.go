package main

import (
	"flag"
	"fmt"
	"os"
	//"log"
	"path/filepath"

	//"database/sql"
	//"github.com/coopernurse/gorp"
	//_ "github.com/mattn/go-sqlite3"
)

var (
	web bool
	spider bool
	//dbMap *gorp.DbMap
)

func init() {
	flag.BoolVar(&web, "w", false, "Run web site")
	flag.BoolVar(&web, "web", false, "Run web site")

	flag.BoolVar(&spider, "s", false, "Run spider")
	flag.BoolVar(&spider, "spider", false, "Run spider")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\nOptions:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	/*
	dbMap := initDb()
	defer dbMap.Db.Close()
	*/
	if web {
		fmt.Println(Movie{})
	}
	/*
	if spider {
		RunSpider()
	}
	*/
}

/*
func initDb() *gorp.DbMap {
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

	return dbMap
}
*/
