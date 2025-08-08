package main

import (
	"flag"
	"fmt"
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
	cliFormat := flag.Bool("cli", false, "show story in terminal")

	flag.Parse()

	story, err := parse.LoadStory(*storyFile)
	if err != nil {
		panic(err)
	}

	if !*cliFormat {
		tmpl := template.Must(template.ParseFiles("templates/page.html"))

		mux := CreateMux(story, tmpl)

		println("Running on  http://localhost:8080")
		err = http.ListenAndServe(":8080", mux)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		chosenArc := "intro"
		fmt.Printf("\n---\n%v\n\n", story["intro"].Title)
		for _, line := range story["intro"].Story {
			println(line)
		}
		for index, option := range story["intro"].Options {
			fmt.Printf("\nOption %v: %v", index+1, option.Text)
		}

		println("\n")

		for {
			fmt.Printf("\nChoice: ")
			var input int
			_, err := fmt.Scan(&input)
			if err != nil {
				return
			}

			if input < 1 || input > len(story["intro"].Options) {
				continue
			}

			chosenArc = story[chosenArc].Options[input-1].Arc

			fmt.Printf("\n---\n%v\n\n", story[chosenArc].Title)
			for _, line := range story[chosenArc].Story {
				println(line)
			}

			if chosenArc == "home" {
				return
			}

			for index, option := range story[chosenArc].Options {
				fmt.Printf("\nOption %v: %v\t", index+1, option.Text)
			}
		}
	}

}
