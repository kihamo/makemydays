package app

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gocraft/web"
)

type Context struct {
}

func (c *Context) RootPage(w web.ResponseWriter, r *web.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, GetRecommendation())
}

func (c *Context) NotFound(w web.ResponseWriter, r *web.Request) {
	w.WriteHeader(http.StatusNotFound)

	t, _ := template.ParseFiles("templates/404.html")
	t.Execute(w, nil)
}

func (c *Context) Error(w web.ResponseWriter, r *web.Request, err interface{}) {
	w.WriteHeader(http.StatusInternalServerError)

	t, _ := template.ParseFiles("templates/500.html")
	t.Execute(w, err)
}

func (c *Context) ApiRecomendations(w web.ResponseWriter, r *web.Request) {
	jsonContent, err := json.Marshal(GetRecommendation())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonContent)
}

func RunServer(addr string) {
	router := web.New(Context{})

	// middleware
	router.
		Middleware(web.LoggerMiddleware).
		Middleware(web.StaticMiddleware("public")).

		// TODO: debug obly
		Middleware(web.ShowErrorsMiddleware)

	// routes
	router.
		NotFound((*Context).NotFound).
		Error((*Context).Error).
		Get("/", (*Context).RootPage)

	routerApi := router.Subrouter(Context{}, "/api/v1")
	routerApi.
		Get("/recomendations", (*Context).ApiRecomendations)

	if err := http.ListenAndServe(addr, router); err != nil {
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
