package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const webPort = "4002"

type Config struct{}

func main() {
	app := Config{}

	fmt.Println("Starting front end service on port 4002")

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
