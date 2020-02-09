package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Body struct {
	Token        string
	Team_id      string
	Api_app_id   string
	Event        Event
	Type         string
	Event_id     string
	Event_time   int
	Authed_users []string
	Challenge    string
}

type Event struct {
	Type     string
	User     string
	Item     Item
	Reaction string
	Event_ts string
}

type Item struct {
	Type    string
	Channel string
	Ts      string
}

func Listen(port int) {
	portString := ":" + strconv.Itoa(port)

	http.HandleFunc("/", hundle)
	http.ListenAndServe(portString, nil)
}

func hundle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var body Body
		if err := decoder.Decode(&body); err != nil {
			log.Fatal(err)
		}

		if body.Type == "url_verification" {
			fmt.Fprint(w, body.Challenge)
			return
		}

		if body.Type == "event_callback" {
			if body.Event.Type != "reaction_added" {
				return
			}

			if body.Event.Reaction != "jp" && body.Event.Reaction != "us" {
				return
			}

			findMessage(body.Event.Item.Channel, body.Event.Item.Ts)

			fmt.Println(body)
		}
	}
}
