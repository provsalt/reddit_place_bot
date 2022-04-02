package main

import (
	"github.com/sirupsen/logrus"
	"image/png"
	"net/http"
	"os"
	"reddit_place_bot/placer"
)

func main() {
	logger := logrus.New()
	data, err := os.Open("image.png")
	if err != nil {
		logger.Fatal(err)
	}
	m, err := png.Decode(data)
	if err != nil {
		logger.Fatal(err)
	}
	client := http.Client{}
	placer.New(m, client, "")
}
