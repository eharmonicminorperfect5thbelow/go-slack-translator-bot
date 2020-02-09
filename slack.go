package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Messages struct {
	Ok       bool
	Messages []Message
}

type Message struct {
	Type    string
	Subtype string
	Ts      string
	User    string
	Text    string
}

func postMessage(text string) {
	return
}

func findMessage(channel string, ts string) {
	req, err := http.NewRequest("GET", "https://slack.com/api/conversations.history", nil)

	params := req.URL.Query()
	params.Add("token", config.SlackAccessToken)
	params.Add("channel", channel)
	params.Add("latest", ts[:len(ts)-1]+"1")
	params.Add("oldest", strings.Split(ts, ".")[0])
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

	var messages Messages
	if err := json.Unmarshal(bytes, &messages); err != nil {
		log.Fatal(err)
	}

	if !messages.Ok {
		return
	}

	if len(messages.Messages) != 1 {
		return
	}

	fmt.Println(messages.Messages[0].Text)

	return
}

func findReply() {
	return
}

func checkExistence() {
	return
}
