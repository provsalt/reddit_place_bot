package canvas

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
)

var redditPlace = url.URL{Scheme: "wss", Host: "gql-realtime-2.reddit.com", Path: "query/"}

func GetCanvas(bearer string) {
	header := http.Header{}
	header.Set("Origin", "https://hot-potato.reddit.com")
	ws, _, err := websocket.DefaultDialer.Dial(redditPlace.String(), header)
	if err != nil {
		panic(err)
	}
	defer ws.Close()
	message := map[string]interface{}{
		"type": "connection_init",
		"payload": map[string]interface{}{
			"Authorization: Bearer": header,
		},
	}
	json2, _ := json.Marshal(message)
	ws.WriteJSON(json2)

}
