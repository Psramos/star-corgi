package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

var astrologyUrl string = "https://www.astroyogi.com/horoscopes/daily/%s-free-horoscope.aspx"

var slackUrl string = "YOUR_SLACK_URL"

type SlackRequestBody struct {
	Text string `json:"text"`
}

func postMessage(text string) {
	body, _ := json.Marshal(SlackRequestBody{Text: text})
	request, err := http.NewRequest(http.MethodPost, slackUrl, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println(err)

		return
	}

	request.Header.Add("Content-type", "application/json")

	client := http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println(err)

		return
	}

	buf := new(bytes.Buffer)

	buf.ReadFrom(resp.Body)

	if buf.String() != "ok" {
		fmt.Println(buf.String())
	}
}
func guidance(horoscope string)  {
	client := http.Client{}
	res, _ := client.Get(fmt.Sprintf(astrologyUrl, horoscope))

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#today1").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		text = strings.Replace(text, "\n", "", 10)
		text = strings.Replace(text, "Click here for a more personalised reading", "", 1)

		postMessage(horoscope)
		postMessage(text)
	})
}
func main() {
	guidance("aries")
}