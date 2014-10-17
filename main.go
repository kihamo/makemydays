package main

import (
	"flag"
	"fmt"
	"os"
	//"log"
	"path/filepath"

	"github.com/kihamo/makemydays/makemydays"

	//"database/sql"
	//"github.com/coopernurse/gorp"
	//_ "github.com/mattn/go-sqlite3"
)

var (
	web bool
	spider bool
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

	if web {
		makemydays.RunServer()
	}

	if spider {
		makemydays.RunSpider()
	}
}
