package main

import (
	"log"
	"net/http"
)

func (app *Config) HandleHTML(w http.ResponseWriter, r *http.Request) {
	app.ShowMessage = true
	app.ShowTable = false
	app.Message = "Welcome to the CRUD application!"
	app.render(w, "test.page.gohtml", app)
}

func (app *Config) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	tempRecord := Record{}

	for key, value := range r.Form {
		if key != "submit" {
			if key == "rid" {
				tempRecord.ID = value[0]
			} else if key == "org" {
				tempRecord.Name = value[0]
			} else if key == "idea" {
				tempRecord.Idea = value[0]
			} else if key == "len" {
				tempRecord.Duration = value[0]
			} else if key == "diff" {
				tempRecord.Difficulty = value[0]
			} else {
				tempRecord.Progress = value[0]
			}
		}
	}

	app.Table = append(app.Table, tempRecord)

	app.ShowMessage = true
	app.ShowTable = false
	app.Message = "The record has been created!"
	app.render(w, "test.page.gohtml", app)
}

func (app *Config) Read(w http.ResponseWriter, r *http.Request) {

	app.ShowMessage = true
	app.ShowTable = true
	app.Message = "The table has been generated!"
	app.render(w, "test.page.gohtml", app)
}

func (app *Config) Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	tempRecord := Record{}

	for key, value := range r.Form {
		if key != "submit" {
			if key == "rid" {
				tempRecord.ID = value[0]
			} else if key == "org" {
				tempRecord.Name = value[0]
			} else if key == "idea" {
				tempRecord.Idea = value[0]
			} else if key == "len" {
				tempRecord.Duration = value[0]
			} else if key == "diff" {
				tempRecord.Difficulty = value[0]
			} else {
				tempRecord.Progress = value[0]
			}
		}
	}

	for index, item := range app.Table {
		if item.ID == tempRecord.ID {
			app.Table = append(app.Table[:index], app.Table[index+1:]...)
			break
		}
	}

	app.Table = append(app.Table, tempRecord)

	app.ShowMessage = true
	app.ShowTable = false
	app.Message = "The table has been updated!"
	app.render(w, "test.page.gohtml", app)
}

func (app *Config) Delete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	tempRecord := Record{}

	for key, value := range r.Form {
		if key != "submit" {
			if key == "rid" {
				tempRecord.ID = value[0]
			} else if key == "org" {
				tempRecord.Name = value[0]
			} else if key == "idea" {
				tempRecord.Idea = value[0]
			} else if key == "len" {
				tempRecord.Duration = value[0]
			} else if key == "diff" {
				tempRecord.Difficulty = value[0]
			} else {
				tempRecord.Progress = value[0]
			}
		}
	}

	for index, item := range app.Table {
		if item.ID == tempRecord.ID {
			app.Table = append(app.Table[:index], app.Table[index+1:]...)
			break
		}
	}

	app.ShowMessage = true
	app.ShowTable = false
	app.Message = "The record has been deleted!"
	app.render(w, "test.page.gohtml", app)
}
