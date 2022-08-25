package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

type HomePageData struct {
	Name   string
	IsLong bool
	Numbers []int
}

func main() {
	r := chi.NewRouter()
	r.Get("/home/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		data := HomePageData{
			Name: name,
			IsLong: len(name) > 4,
			Numbers: []int{1, 2, 4},
		}
		t, err := template.ParseFiles("pages/home.html")
		if err != nil {
			fmt.Println("err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = t.Execute(w, data); err != nil {
			fmt.Println("err:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":9000", r)
}
