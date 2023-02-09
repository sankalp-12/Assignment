package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const webPort = "4040"

type Config struct {
	Table       []Record
	Message     string
	ShowTable   bool
	ShowMessage bool
}

type Record struct {
	ID         string
	Name       string
	Idea       string
	Duration   string
	Difficulty string
	Progress   string
}

func main() {
	app := Config{}

	// Populate the Table with some predefined entries
	app.Table = append(app.Table, Record{ID: "S. No.", Name: "Name of the Organization", Idea: "Project", Duration: "Duration", Difficulty: "Difficulty", Progress: "Progress"})
	app.Table = append(app.Table, Record{ID: "1", Name: "MetaCall", Idea: "Builder", Duration: "175 hours", Difficulty: "Medium", Progress: "Zero contributions"})
	app.Table = append(app.Table, Record{ID: "2", Name: "JBoss Community", Idea: "Kroxylicious", Duration: "175 hours", Difficulty: "Medium", Progress: "Zero contributions"})

	fmt.Println("Starting front end service on port 4040")

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

//go:embed template
var templateFS embed.FS

func (app *Config) render(w http.ResponseWriter, t string, data any) {

	partials := []string{
		"template/base.layout.gohtml",
		"template/header.partial.gohtml",
		"template/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("template/%s", t))

	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if data == nil {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
