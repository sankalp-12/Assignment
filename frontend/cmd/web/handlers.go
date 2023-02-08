package main

import (
	"log"
	"net/http"
)

type Data struct {
	Site  string
	YesNo bool
}

func (app *Config) HandleHTML(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *Config) HandleURL(w http.ResponseWriter, r *http.Request) {
	chanURL := make(chan string)
	len := 0

	var output string
	var yesno bool
	var dataPassed []Data

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		var urls []string

		for key, value := range r.Form {
			if key != "submit" {
				urls = append(urls, value[0])
				len = len + 1
			}
		}

		for _, webpage := range urls {
			// log.Println(webpage)
			go getWebpage(webpage, chanURL)
		}

		for {
			msg := <-chanURL
			// log.Println("Input Processed!")
			output = msg
			if output[0] == 'Y' {
				yesno = true
			} else {
				yesno = false
			}
			tempData := Data{
				Site:  output,
				YesNo: yesno,
			}
			dataPassed = append(dataPassed, tempData)
			len = len - 1
			if len == 0 {
				close(chanURL)
				break
			}
		}
	}

	app.render(w, "test.page.gohtml", dataPassed)
}

func getWebpage(webpage string, chanURL chan string) {
	if _, err := http.Get(webpage); err != nil {
		if webpage == "" {
			chanURL <- "Oops! You did not enter any website!"
			// log.Println("Output Processed --> NO!")
		} else {
			chanURL <- "Nope, the website " + webpage + " is down at the moment!"
			// log.Println("Output Processed --> NO!")
		}
	} else {
		chanURL <- "Yes, the website " + webpage + " is up and running!"
		// log.Println("Output Processed --> YES!")
	}
}
