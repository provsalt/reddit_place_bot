package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"image/png"
	"net/http"
	"os"
	config2 "reddit_place_bot/config"
	"reddit_place_bot/httputils"
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
	cfg, _ := os.ReadFile("config.json")
	config := config2.Config{}
	err = json.Unmarshal(cfg, &config)
	if err != nil {
		logger.Fatal(err)
	}

	for i := 0; i < len(config.Accounts); i++ {
		i := i
		go func() {
			client := http.Client{}
			httputils.RefreshBearer(config.Accounts[i].Username, config.Accounts[i].Password, config.Accounts[i].ClientID, config.Accounts[i].ClientSecret)
			placer.New(m, client, "")
		}()
	}
}
