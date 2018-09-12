package main

import (
	"html/template"
	"log"
	"net/http"

	"./jsonParser"
	s "./structs"
	//"fmt"
	"github.com/gorilla/mux"
)

func main() {
	var stories map[string]s.Story
	gotStories := make(chan bool)

	go jsonParser.GetJson(&stories, gotStories)
	<-gotStories

	startingPoint := "intro"
	currentStory := stories[startingPoint]

	tmpl := template.Must(template.ParseFiles("src/templates/index.html"))

	//TODO: refactor server into different file
	r := mux.NewRouter()
	s := r.PathPrefix("/").Subrouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := currentStory
		tmpl.Execute(w, data)
	})

	//TODO: Work on adding styling to the HTML template
	s.HandleFunc("/{arc}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		category := vars["arc"]

		//TODO: add error handling to this in case the category is not found in stories
		data := stories[category]
		tmpl.Execute(w, data)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}

/*
 *TODO: Refactor the server for this using server as struct and extending methods from it
 *
 *https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
 */
