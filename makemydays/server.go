package makemydays

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	ADDR = ":9001"
)

func RunServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/recomendations", IndexHandler).Methods("GET")

	http.Handle("/", r)
	if err := http.ListenAndServe(ADDR, nil); err != nil {
		log.Fatalln("failed to start server", err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	db := NewDatabase()
	defer db.Close()

	recommendation := &Recommendation{}

	movie := &Movie{}
	if err := db.Limit(1).Order("random()").Find(movie).Error; err == nil {
		recommendation.Movie = movie
	}

	song := &Song{}
	if err := db.Limit(1).Order("random()").Find(song).Error; err == nil {
		recommendation.Song = song
	}

	book := &Book{}
	if err := db.Limit(1).Order("random()").Find(book).Error; err == nil {
		recommendation.Book = book
	}

	word := &Word{}
	if err := db.Limit(1).Order("random()").Find(word).Error; err == nil {
		recommendation.Word = word
	}

	task := &Task{}
	if err := db.Limit(1).Order("random()").Find(task).Error; err == nil {
		recommendation.Task = task
	}

	food := &Food{}
	if err := db.Limit(1).Order("random()").Find(food).Error; err == nil {
		recommendation.Food = food
	}

	jsonContent, err := json.Marshal(recommendation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonContent)
}
