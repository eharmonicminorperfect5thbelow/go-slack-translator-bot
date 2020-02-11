package bot

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func translate(text string, from string, to string) string {
	req, err := http.NewRequest("GET", config.TranslationURL, nil)

	params := req.URL.Query()
	params.Add("text", text)
	params.Add("from", from)
	params.Add("to", to)
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

	return string(bytes)
}

func translateMessage(text string, from string, to string) string {
	re := regexp.MustCompile("<@.{9}>|<!channel>|<!here>")
	mention := re.FindAllStringSubmatch(text, -1)
	splitMessage := re.Split(text, -1)
	translatedMessage := make([]string, len(splitMessage))
	joinedMessage := translate(splitMessage[0], from, to)

	for i, s := range splitMessage {
		translatedMessage[i] = translate(s, from, to)
	}

	for i := 0; i < len(mention); i++ {
		joinedMessage += mention[i][0] + translatedMessage[i+1]
	}

	return joinedMessage
}
