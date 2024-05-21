package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("base.html"))

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := struct {
			Title        template.HTML
			BusinessName string
			Slogan       string
		}{
			Title:        template.HTML("Ex3_week3 | AI & GPT"),
			BusinessName: "Business,",
			Slogan:       "we get things done!",
		}
		err := templates.ExecuteTemplate(w, "base.html", &b)
		if err != nil {
			fmt.Printf("index: couldn't parse template: %v\n", err)
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func main() {
	http.Handle("/", index())
	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
