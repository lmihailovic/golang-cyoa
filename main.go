package main

import (
	"flag"
	"golang-cyoa/parse"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func CreateMux(story parse.Story, tmpl *template.Template) http.Handler {
	mux := http.NewServeMux()

	for name, chapter := range story {
		path := "/" + strings.ReplaceAll(name, " ", "-")
		log.Println("Adding route for", path)

		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			tmpl.Execute(w, chapter)
		})

	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/intro", http.StatusFound)
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	return mux
}

func main() {
	storyFile := flag.String("f", "gopher.json", "path to story file")

	flag.Parse()

	story, err := parse.LoadStory(*storyFile)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("templates/page.html"))

	mux := CreateMux(story, tmpl)

	println("Running on  http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
