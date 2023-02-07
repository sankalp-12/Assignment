package main

import (
	"encoding/csv"
	"net/http"
)

type Record struct {
	Row []string
}

func (app *Config) HandleHTML(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *Config) HandleCSV(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tempFile, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pointerCSV := csv.NewReader(tempFile)
		dataCSV, err := pointerCSV.ReadAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sliceOfStruct := Parse(dataCSV)
		app.render(w, "test.page.gohtml", sliceOfStruct)
	}
}

func Parse(dataCSV [][]string) []Record {
	var sliceOfStruct []Record
	for _, line := range dataCSV {
		var tempRow Record

		tempRow.Row = append(tempRow.Row, line...)

		sliceOfStruct = append(sliceOfStruct, tempRow)
	}
	return sliceOfStruct
}
