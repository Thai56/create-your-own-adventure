package main

import (
	"fmt"
	//"html/template"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"path/filepath"
)

type OptionType struct {
	Text string
	Arc  string
}

type Story struct {
	Title   string
	Story   string
	Options []OptionType
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!!!")
	w.Write([]byte("Hello World!!!"))
}

func main() {
	var stories map[string]Story
	gotStories := make(chan bool)

	go getJson(&stories, gotStories)
	<-gotStories

	startingPoint := "intro"
	currentStory := stories[startingPoint]

	fmt.Println("Current Story ", currentStory)
	//TODO: Set up basic web server
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
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
