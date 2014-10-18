package makemydays

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	API_URL = "http://makemydays.me/allapi/"
)

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
	movieValue := strings.TrimSpace(jsonObject["Filmapi"].(string))
	movieYear, _ := strconv.ParseInt(movieValue[len(movieValue)-4:len(movieValue)], 10, 64)

	movie := Movie{
		Title: strings.TrimSpace(movieValue[:len(movieValue)-4]),
		Year: movieYear,
	}
	db.Save(&movie)

	// song
	song := Song{
		Title: strings.TrimSpace(jsonObject["Musicapi"].(string)),
	}
	db.Save(&song)

	// word
	word := Word{
		Word: strings.TrimSpace(jsonObject["Wordapi"].(string)),
	}
	db.Save(&word)

	// book
	book := Book{
		Title: strings.TrimSpace(jsonObject["Bookapi"].(string)),
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
