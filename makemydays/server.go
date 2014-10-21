package makemydays

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer(addr string) {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))

	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/recomendations", RecommendationsHandler).Methods("GET")

	http.Handle("/", r)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln("failed to start server", err)
	}
}

func GetRecommendation() *Recommendation {
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

	return recommendation
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, GetRecommendation())
}

func RecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	jsonContent, err := json.Marshal(GetRecommendation())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonContent)
}
