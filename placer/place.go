package placer

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
)

type Place struct {
	i      image.Image
	client http.Client
}

func New(image image.Image, client http.Client, token string) *Place {
	return &Place{
		i:      image,
		client: client,
	}
}

func (p Place) imageToRGBA() *image.RGBA {
	bounds := p.i.Bounds()
	rgba := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, p.i.At(x, y))
		}
	}
	return rgba
}

func (p *Place) Ticker() {

}

func (p Place) PlacePixel(x, y, c int) {
	bigBodyTemplate := `{"operationName":"setPixel","variables":{"input":{"actionName":"r/replace:set_pixel","PixelMessageData":{"coordinate":{"x":%d,"y":%d},"colorIndex":14,"canvasIndex":0}}},"query":"mutation setPixel($input: ActInput!) {\n act(input: $input) {\n data {\n ... on BasicMessage {\n id\n data {\n ... on GetUserCooldownResponseMessageData {\n nextAvailablePixelTimestamp\n __typename\n }\n ... on SetPixelResponseMessageData {\n timestamp\n __typename\n }\n __typename\n }\n __typename\n }\n __typename\n }\n __typename\n }\n}\n"}`
	requestBody := fmt.Sprintf(bigBodyTemplate, x, y)

	req, err := http.NewRequest("POST", "https://gql-realtime-2.reddit.com/query", bytes.NewBuffer([]byte(requestBody)))

	req.Header.Add("accept", "*/*")
	req.Header.Add("apollographql-client-name", "mona-lisa")
	req.Header.Add("apollographql-client-version", "0.0.1")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", myBearerToken))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")

	resp, err := p.client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

}

func (p Place) PlaceImage() {
	rgba := p.imageToRGBA()
	for y := 0; y < rgba.Bounds().Max.Y; y++ {
		for x := 0; x < rgba.Bounds().Max.X; x++ {
			p.PlacePixel(x, y, int(rgba.At(x, y).(color.RGBA).R))
		}
	}
}

func closestColor(target color.Color) int {
	r, g, b, a := target.RGBA()
	for i, i2 := range  {
		
	}
}
