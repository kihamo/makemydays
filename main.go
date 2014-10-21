package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kihamo/makemydays/makemydays"
)

var (
	web bool
	spider bool
	addr string
)

func init() {
	flag.BoolVar(&web, "w", false, "Run web site")
	flag.BoolVar(&web, "web", false, "Run web site")

	flag.BoolVar(&spider, "s", false, "Run spider")
	flag.BoolVar(&spider, "spider", false, "Run spider")

	flag.StringVar(&addr, "a", ":9001", "Server address listen")
	flag.StringVar(&addr, "addr", ":9001", "Server address listen")

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
		makemydays.RunServer(addr)
	}

	if spider {
		makemydays.RunSpider()
	}
}
