package makemydays

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

	"github.com/jinzhu/gorm"
)

const (
	//API_URL = "http://makemydays.me/allapi/"
	API_URL = "http://kvartal918.kihamo.ru/make.json"
	MAX_WORKERS = runtime.NumCPU()
	MAX_ITERATIONS = 100
)

var db *gorm.DB

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

	req.Header.Add("User-Agent", fmt.Sprintf("Go-GO Krankus! (request D: %d)", requestId))

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

func saveData(object map[string]interface{}) {
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
	db.Save(&movie)

	// song
	originalValue = object["Musicapi"].(string)
	songData := strings.Split(originalValue, " – ")
	song := Song{
		OriginalValue: originalValue,
		Title: strings.TrimSpace(songData[1]),
		Author: strings.TrimSpace(songData[0]),
	}
	db.Save(&song)

	// book
	originalValue = object["Bookapi"].(string)
	bookData := strings.Split(originalValue, " — ")
	book := Book{
		OriginalValue: originalValue,
		Title: TrimFuncQuote(bookData[0]),
		Author: strings.TrimSpace(bookData[1]),
	}
	db.Save(&book)

	// word
	word := Word{
		Word: strings.TrimSpace(object["Wordapi"].(string)),
	}
	db.Save(&word)

	// task
	task := Task{
		Title: strings.TrimSpace(object["Taskapi"].(string)),
	}
	db.Save(&task)

	// food
	food := Food{
		Title: strings.TrimSpace(object["Foodapi"].(string)),
	}
	db.Save(&food)

	fmt.Printf("Movie: %v; Song: %v; Book: %v; Word: %v; Task: %v; Food: %v;\n", movie, song, book, word, task, food)
}

func spider(done chan<- struct{}, requestId int) {
	object, err := loadDataFromApi(requestId)

	if err != nil {
		log.Println("Load data from api failed. Error:", err)
	} else {
		saveData(object)
	}
	done <- struct{}{}
}

func RunSpider() {
	doneIterations := 0
	doneWorkers := 0

	runtime.GOMAXPROCS(runtime.NumCPU())
	done := make(chan struct{}, MAX_WORKERS)

	db = NewDatabase()
	defer db.Close()

	for i := 0; i < MAX_WORKERS; {
		i++

		doneIterations++
		go spider(done, doneIterations)
	}

	for doneIterations < MAX_ITERATIONS || doneWorkers != MAX_WORKERS {
		select {
		case <-done:
			doneWorkers++

			if doneIterations < MAX_ITERATIONS {
				doneIterations++
				doneWorkers--

				go spider(done, doneIterations)
			}
		}
	}
}
