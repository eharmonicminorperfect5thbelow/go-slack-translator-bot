package bot

import (
	"encoding/json"
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

type Result struct {
	Ok bool
}

func postMessage(text string, channel string) {
	req, err := http.NewRequest("GET", "https://slack.com/api/chat.postMessage", nil)

	params := req.URL.Query()
	params.Add("token", config.SlackAccessToken)
	params.Add("text", text)
	params.Add("channel", channel)
	req.URL.RawQuery = params.Encode()

	timeout := time.Duration(10 * time.Second)
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

	var result Result
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Fatal(err)
	}

	if !result.Ok {
		log.Fatal("failed to post message")
	}
}

func findMessage(channel string, ts string) string {
	req, err := http.NewRequest("GET", "https://slack.com/api/conversations.history", nil)

	params := req.URL.Query()
	params.Add("token", config.SlackAccessToken)
	params.Add("channel", channel)
	params.Add("latest", ts[:len(ts)-1]+"1")
	params.Add("oldest", strings.Split(ts, ".")[0])
	req.URL.RawQuery = params.Encode()

	timeout := time.Duration(10 * time.Second)
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
		log.Fatal("failed to find message")
	}

	if len(messages.Messages) != 1 {
		log.Fatal("message not found")
	}

	return messages.Messages[0].Text
}

func findReply() {
	return
}
