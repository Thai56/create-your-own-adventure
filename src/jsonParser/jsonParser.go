package jsonParser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	s "../structs"
)

func GetJson(stories *map[string]s.Story, finished chan bool) {
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
