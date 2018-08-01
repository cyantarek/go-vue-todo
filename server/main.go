package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"database/sql"
	"log"
	"html/template"
	"github.com/go-chi/chi/middleware"
)

func main() {
	DbSetup()
	r := chi.NewMux()
	r.Use(middleware.DefaultLogger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.New("index.html").Delims("[[", "]]").ParseFiles("index.html")
		if err != nil {
			log.Fatalln(err)
		}
		tpl.Execute(w, nil)
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Use(JsonResponse)
		r.Get("/", GetTasks)
		r.Post("/", PostTasks)
		r.Delete("/{id}", DeleteTasks)
	})

	http.ListenAndServe(":8076", r)
}

func DbSetup() {
	var err error
	Db, err = sql.Open("sqlite3", "storage.db")

	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec("create table if not exists tasks(id integer not null primary key autoincrement, name varchar not null)")

	if err != nil {
		log.Fatal(err)
	}
}
