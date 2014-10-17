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

	dbMap := GetDatabase()
	defer dbMap.Db.Close()
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	movie := Movie{}
	dbMap := GetDatabase()

	err := dbMap.SelectOne(&movie, "SELECT * FROM movies ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	random := Recommendation{movie, Song{}, Word{}, Book{}, Task{}, Food{}}

	jsonContent, err := json.Marshal(random)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonContent)
}
