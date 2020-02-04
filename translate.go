package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Translate(text string, from string, to string) string {
	req, err := http.NewRequest("GET", config.TranslationURL, nil)

	params := req.URL.Query()
	params.Add("text", text)
	params.Add("from", from)
	params.Add("to", to)
	req.URL.RawQuery = params.Encode()

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
