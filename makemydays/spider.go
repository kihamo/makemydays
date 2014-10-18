package makemydays

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	API_URL = "http://makemydays.me/allapi/"
)

func TrimFuncQuote(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
			return r == '«' || r == '»'
		});
}

func RunSpider() {
	req, err := http.NewRequest("GET", API_URL, nil)
	if req == nil {
		log.Fatalln("Reguest makemydays api failed", err)
	}

	req.Header.Add("User-Agent", "Bad bot from RuRu O_o Go-GO Krankus!")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("Reguest makemydays api failed", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Reguest response makemydays api failed", err)
	}

	var object interface{}
	if err := json.Unmarshal(body, &object); err != nil {
		log.Fatalln("Reguest parse json makemydays api failed", err)
	}

	jsonObject := object.(map[string]interface{})

	db := NewDatabase()
	defer db.Close()

	// movie
	re := regexp.MustCompile(`(?:([^\p{Cyrillic}]+)\p{Zs})?(\p{Cyrillic}.*)?\p{Zs}([\d]+)`)
	movieValue := strings.TrimSpace(jsonObject["Filmapi"].(string))
	movieData := re.FindAllStringSubmatch(movieValue, -1)

	if movieData == nil {
		log.Fatalln("Error parse movie string:", movieValue)
	}

	movieYear, err := strconv.ParseInt(movieData[0][3], 10, 64)
	if err != nil {
		log.Fatalln("Convert year to integer failed", err)
	}

	movie := Movie{
		Title: movieData[0][1],
		TitleRus: movieData[0][2],
		Year: movieYear,
	}
	db.Save(&movie)

	// song
	songData := strings.Split(jsonObject["Musicapi"].(string), " – ")
	song := Song{
		Title: strings.TrimSpace(songData[1]),
		Author: strings.TrimSpace(songData[0]),
	}
	db.Save(&song)

	// word
	word := Word{
		Word: strings.TrimSpace(jsonObject["Wordapi"].(string)),
	}
	db.Save(&word)

	// book
	bookData := strings.Split(jsonObject["Bookapi"].(string), " — ")
	book := Book{
		Title: TrimFuncQuote(bookData[0]),
		Author: strings.TrimSpace(bookData[1]),
	}
	db.Save(&book)

	// task
	task := Task{
		Title: strings.TrimSpace(jsonObject["Taskapi"].(string)),
	}
	db.Save(&task)

	// food
	food := Food{
		Title: strings.TrimSpace(jsonObject["Foodapi"].(string)),
	}
	db.Save(&food)
}
