package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type OptionType struct {
	Text string
	Arc  string
}

type Story struct {
	Title   string
	Story   []string
	Options []OptionType
}

func main() {
	var stories map[string]Story
	gotStories := make(chan bool)

	go getJson(&stories, gotStories)
	<-gotStories

	startingPoint := "intro"
	currentStory := stories[startingPoint]

	tmpl := template.Must(template.ParseFiles("src/index.html"))

	r := mux.NewRouter()
	s := r.PathPrefix("/").Subrouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := currentStory
		tmpl.Execute(w, data)
	})

	s.HandleFunc("/{arc}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		category := vars["arc"]
		fmt.Println("category", category)

		data := stories[category]
		tmpl.Execute(w, data)
	})

	//s.Path("/{key}/").
	//HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//category := vars["key"]
	//fmt.Println("category", category)
	//fmt.Println("reaching here")
	//})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
	//TODO: work on putting story as html on browser
}

func getJson(stories *map[string]Story, finished chan bool) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Problem getting working directory: %v\n", err)
	}

	jsonUnparsed, err := ioutil.ReadFile(pwd + "/gopher.json")
	if err != nil {
		fmt.Printf("There was a problem reading the file: %v \n", err)
	}

	json.Unmarshal(jsonUnparsed, stories)

	finished <- true
}
