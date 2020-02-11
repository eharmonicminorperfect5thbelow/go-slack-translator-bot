package bot

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	SlackAccessToken string
	TranslationURL   string
}

var config Config

func loadConfig(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatal(err)
	}
}
