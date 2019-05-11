package bot

import (
	"bytes"
	"log"
	"net/http"
	"os"
)

var url string

func init() {
	slack := os.Getenv("SLACK_URL")
	if slack == "" {
		log.Fatal("$SLACK_URL must be set")
	}
	url = slack
}

func SendMsg(jsonStr []byte) {
	//curl -X POST -H 'Content-type: application/json' --data '{"text":"Hello, World!"}' URL

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
