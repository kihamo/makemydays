package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	API_URL = "http://makemydays.me/allapi/"
	MAX_WORKERS = 4
	MAX_ITERATIONS = 100
)

func TrimFuncQuote(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
			return r == '«' || r == '»'
		});
}

func loadDataFromApi(requestId int) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", API_URL, nil)
	if req == nil {
		return nil, err
	}

	req.Header.Add("User-Agent", fmt.Sprintf("Go-GO Krankus! (request ID: %d)", requestId))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var object interface{}
	if err := json.Unmarshal(body, &object); err != nil {
		return nil, err
	}

	return object.(map[string]interface{}), nil
}

func saveData(results chan<- interface{}, object map[string]interface{}) {
	var originalValue string

	// movie
	re := regexp.MustCompile(`(?:([^\p{Cyrillic}]+)\p{Zs})?(\p{Cyrillic}.*)?\p{Zs}([\d]+)`)
	originalValue = strings.TrimSpace(object["Filmapi"].(string))
	movieData := re.FindAllStringSubmatch(originalValue, -1)

	if movieData == nil {
		log.Fatalln("Error parse movie string:", originalValue)
	}

	movieYear, err := strconv.ParseInt(movieData[0][3], 10, 64)
	if err != nil {
		log.Fatalln("Convert year to integer failed", err)
	}

	movie := Movie{
		OriginalValue: originalValue,
		Title: movieData[0][1],
		TitleRus: movieData[0][2],
		Year: movieYear,
	}
	results <- movie

	// song
	originalValue = object["Musicapi"].(string)
	songData := strings.Split(originalValue, " – ")
	song := Song{
		OriginalValue: originalValue,
		Title: strings.TrimSpace(songData[1]),
		Author: strings.TrimSpace(songData[0]),
	}
	results <- song

	// book
	originalValue = object["Bookapi"].(string)
	bookData := strings.Split(originalValue, " — ")
	book := Book{
		OriginalValue: originalValue,
		Title: TrimFuncQuote(bookData[0]),
		Author: strings.TrimSpace(bookData[1]),
	}
	results <- book

	// word
	word := Word{
		Word: strings.TrimSpace(object["Wordapi"].(string)),
	}
	results <- word

	// task
	task := Task{
		Title: strings.TrimSpace(object["Taskapi"].(string)),
	}
	results <- task

	// food
	food := Food{
		Title: strings.TrimSpace(object["Foodapi"].(string)),
	}
	results <- food
}

func spider(results chan<- interface{}, done chan<- struct{}, requestId int) {
	object, err := loadDataFromApi(requestId)

	if err != nil {
		log.Println("Load data from api failed. Error:", err)
	} else {
		saveData(results, object)
	}

	done <- struct{}{}
}

func RunSpider() {
	start := time.Now()
	defer func() {
		log.Printf("Execution time %s", time.Since(start))
	}()

	runIterations := 0
	runWorkers := 0

	runtime.GOMAXPROCS(runtime.NumCPU())

	maxWorkers := MAX_WORKERS
	if MAX_ITERATIONS < maxWorkers {
		maxWorkers = MAX_ITERATIONS
	}

	results := make(chan interface{}, maxWorkers)
	done := make(chan struct{}, maxWorkers)

	db := NewDatabase()
	defer db.Close()

	for runIterations < maxWorkers {
		runIterations++
		runWorkers++

		go spider(results, done, runIterations)
	}

	for runWorkers > 0 {
		select {
		case result := <-results:
			db.Save(result)

		case <-done:
			if runIterations < MAX_ITERATIONS {
				runIterations++
				go spider(results, done, runIterations)
			} else {
				runWorkers--
			}
		}
	}

	// для случая, когда все подпрограммы завершились, но результаты еще не обработались
	for {
		select {
		case result := <-results:
			db.Save(result)

		default:
			return
		}
	}
}
