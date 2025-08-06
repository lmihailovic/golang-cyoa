package main

import (
	"flag"
	"fmt"
	"golang-cyoa/parse"
	"log"
	"net/http"
	"strings"
)

func CreateMux(story parse.Story) http.Handler {
	mux := http.NewServeMux()

	for name, chapter := range story {
		path := "/" + strings.ReplaceAll(name, " ", "-")
		log.Println("Adding route for", path)

		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<h1> %v </h1>", chapter.Title)

			for _, paragraph := range chapter.Story {
				fmt.Fprintf(w, "<p> %v </p>", paragraph)
			}

			for _, option := range chapter.Options {
				fmt.Fprintf(w, "<p> <a href='/%v'>%v</a> </p>", option.Arc, option.Text)
			}
		})

	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusFound)
	})

	return mux
}

func main() {
	storyFile := flag.String("f", "gopher.json", "path to story file")

	story, err := parse.LoadStory(*storyFile)
	if err != nil {
		panic(err)
	}

	mux := CreateMux(story)

	println("Listening on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
